package conf

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQL struct {
	Host     string `json:"host" toml:"host" env:"MYSQL_HOST"`
	Port     int    `json:"port" toml:"port" env:"MYSQL_PORT"`
	DB       string `json:"db" toml:"db" env:"MYSQL_DB"`
	Username string `json:"username" toml:"username" env:"MYSQL_USERNAME"`
	Password string `json:"password" toml:"password" env:"MYSQL_PASSWORD"`

	// 高级参数
	MaxOpenConn int `toml:"max_open_conn" env:"MYSQL_MAX_OPEN_CONN"`
	MaxIdleConn int `toml:"max_idle_conn" env:"MYSQL_MAX_IDLE_CONN"`
	MaxLifeTime int `toml:"max_life_time" env:"MYSQL_MAX_LIFE_TIME"`
	MaxIdleTime int `toml:"max_idle_time" env:"MYSQL_MAX_IDLE_TIME"`

	// 面临并发安全
	lock sync.Mutex
	db   *gorm.DB
}

func NewDefaultMySQL() *MySQL {
	return &MySQL{
		Host:     "127.0.0.1",
		Port:     3306,
		DB:       "mcenter",
		Username: "root",
		Password: "123456",

		MaxOpenConn: 10,
		MaxIdleConn: 10,
		MaxLifeTime: 60,
		MaxIdleTime: 60,
	}
}

func (m *MySQL) Close() error {
	if m.db == nil {
		return nil
	}

	// 没有提供Close方法
	// 没法实现数据库的优雅关闭
	return nil
}

// 获取连接池
func (m *MySQL) GetConnPool() (*sql.DB, error) {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		m.Username, m.Password, m.Host, m.Port, m.DB)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database %s: %w", dsn, err)
	}

	// 对连接池进行配置
	db.SetMaxOpenConns(m.MaxOpenConn)
	db.SetMaxIdleConns(m.MaxIdleConn)
	if m.MaxLifeTime > 0 {
		db.SetConnMaxLifetime(time.Duration(m.MaxLifeTime) * time.Second)
	}
	if m.MaxIdleTime > 0 {
		db.SetConnMaxIdleTime(time.Duration(m.MaxIdleTime) * time.Second)
	}

	// 测试连接
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping database %s: %s", dsn, err.Error())
	}

	return db, nil
}

// 获取gorm
func (m *MySQL) ORM() *gorm.DB {
	m.lock.Lock()
	defer m.lock.Unlock()

	if m.db != nil {
		return m.db
	}

	p, err := m.GetConnPool()
	if err != nil {
		panic(fmt.Errorf("failed to get connection pool: %w", err))
	}

	m.db, err = gorm.Open(mysql.New(mysql.Config{
		Conn: p,
	}), &gorm.Config{
		// 预编译语句 执行任何 SQL 时都创建并缓存预编译语句，可以提高后续的调用速度
		PrepareStmt: true,
		// 跳过默认事务 默认情况下，GORM 会自动开启事务，如果不想开启事务，可以设置 SkipDefaultTransaction 为 true
		// 但是，如果需要使用事务，需要手动开启，对于写操作（创建、更新、删除），为了确保数据的完整性，GORM 会将它们封装在事务内运行。
		// 在不需要事务的场景下，设置为 true 初始化时禁用它，这将获得大约 30%+ 性能提升
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic(err)
	}
	return m.db
}

type Http struct {
	Host string `json:"host" toml:"host" env:"HTTP_HOST"`
	Port int    `json:"port" toml:"port" env:"HTTP_PORT"`
}

func NewDefaultHttp() *Http {
	return &Http{
		Host: "127.0.0.1",
		Port: 8080,
	}
}

func (h *Http) Address() string {
	return fmt.Sprintf("%s:%d", h.Host, h.Port)
}

type Grpc struct {
	Host string `json:"host" toml:"host" env:"GRPC_HOST"`
	Port int    `json:"port" toml:"port" env:"GRPC_PORT"`
}

func NewDefaultGrpc() *Grpc {
	return &Grpc{
		Host: "127.0.0.1",
		Port: 18080,
	}
}

func (g *Grpc) Address() string {
	return fmt.Sprintf("%s:%d", g.Host, g.Port)
}

type MongoDB struct {
	Endpoints []string `toml:"endpoints" env:"MONGO_ENDPOINTS" envSeparator:","`
	UserName  string   `toml:"username" env:"MONGO_USERNAME"`
	Password  string   `toml:"password" env:"MONGO_PASSWORD"`
	Database  string   `toml:"database" env:"MONGO_DATABASE"`
	AuthDB    string   `toml:"auth_db" env:"MONGO_AUTH_DB"`

	client *mongo.Client
	lock   sync.Mutex
}

func NewDefaultMongoDB() *MongoDB {
	return &MongoDB{
		Endpoints: []string{"127.0.0.1:27017"},
		UserName:  "root",
		Password:  "example",
		Database:  "admin",
		AuthDB:    "admin",
	}
}

// 关闭 mongo 客户端
func (m *MongoDB) Close(ctx context.Context) error {
	if m.client == nil {
		return nil
	}
	return m.client.Disconnect(ctx)
}

func (m *MongoDB) GetAuthDB() string {
	if m.AuthDB != "" {
		return m.AuthDB
	}
	return m.Database
}

func (m *MongoDB) GetDB() (*mongo.Database, error) {
	conn, err := m.Client()
	if err != nil {
		return nil, err
	}
	return conn.Database(m.Database), nil
}

func (m *MongoDB) createClient() (*mongo.Client, error) {
	opts := options.Client()

	if m.UserName != "" && m.Password != "" {
		cred := options.Credential{
			AuthSource: m.GetAuthDB(),
		}

		cred.Username = m.UserName
		cred.Password = m.Password
		cred.PasswordSet = true

		opts.SetAuth(cred)
	}

	opts.SetHosts(m.Endpoints)
	opts.SetConnectTimeout(5 * time.Second)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("create mongo client error: %s", err)
	}

	if err = client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("ping mongo server %s error: %s", m.Endpoints, err)
	}

	return client, nil
}

// Client 获取一个全局的mongodb客户端连接
func (m *MongoDB) Client() (*mongo.Client, error) {
	m.lock.Lock()
	defer m.lock.Unlock()

	if m.client != nil {
		return m.client, nil
	}

	conn, err := m.createClient()
	if err != nil {
		return nil, err
	}
	return conn, nil
}

type App struct {
	Name string `json:"name" toml:"name"`
}

func NewDefaultApp() *App {
	return &App{
		Name: "cmdb",
	}
}

type Config struct {
	AppName *App     `json:"app" toml:"app"`
	MySQL   *MySQL   `json:"mysql" toml:"mysql"`
	MongoDB *MongoDB `json:"mongo" toml:"mongo"`
	Http    *Http    `json:"http" toml:"http"`
	Grpc    *Grpc    `json:"grpc" toml:"grpc"`
}

func DefaultConfig() *Config {
	return &Config{
		AppName: NewDefaultApp(),
		MySQL:   NewDefaultMySQL(),
		MongoDB: NewDefaultMongoDB(),
		Http:    NewDefaultHttp(),
		Grpc:    NewDefaultGrpc(),
	}
}

// 将配置文件格式化 Json 文件输出
func (c *Config) String() string {
	d, _ := json.MarshalIndent(c, "", " ")
	return string(d)
}
