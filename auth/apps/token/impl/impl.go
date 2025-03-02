package impl

import (
	"auth/apps/token"
	"auth/apps/user"
	"auth/conf"

	"github.com/infraboard/mcube/ioc"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

// 显示的约束对象实现接口
var _ token.Service = (*impl)(nil)

type impl struct {
	token.UnimplementedRPCServer
	ioc.ObjectImpl
	user user.Service
	col  *mongo.Collection
}

func (i *impl) Init() error {
	db, err := conf.C().MongoDB.GetDB()
	if err != nil {
		panic(err)
	}
	i.col = db.Collection("tokens")
	i.user = ioc.GetController(user.AppName).(user.Service)
	return nil
}

func (i *impl) Name() string {
	return token.AppName
}

func (i *impl) Registry(server *grpc.Server) {
	token.RegisterRPCServer(server, i)
}

func init() {
	ioc.RegistryController(&impl{})
}
