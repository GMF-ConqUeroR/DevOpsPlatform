package impl

import (
	"auth/apps/endpoint"
	"auth/conf"

	"github.com/infraboard/mcube/ioc"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func init() {
	ioc.RegistryController(&impl{})
}

// 显示的约束对象实现接口
var _ endpoint.Service = (*impl)(nil)

type impl struct {
	// 继承GRPC 服务方法模板
	endpoint.UnimplementedRPCServer
	// ioc
	ioc.ObjectImpl
	// 依赖col
	col *mongo.Collection
}

func (i *impl) Name() string {
	return endpoint.AppName
}

func (i *impl) Init() error {
	// 补充Mongo依赖
	db, err := conf.C().MongoDB.GetDB()
	if err != nil {
		panic(err)
	}
	i.col = db.Collection(endpoint.AppName)
	return nil
}

func (i *impl) Registry(server *grpc.Server) {
	endpoint.RegisterRPCServer(server, i)
}
