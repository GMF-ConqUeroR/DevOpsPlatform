package impl

import (
	"auth/apps/role"
	"auth/conf"

	"github.com/infraboard/mcube/ioc"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

var _ role.Service = (*impl)(nil)

type impl struct {
	role.UnimplementedRPCServer
	ioc.ObjectImpl
	col *mongo.Collection
}

func (i *impl) Name() string {
	return role.AppName
}

func (i *impl) Init() error {
	db, err := conf.C().MongoDB.GetDB()
	if err != nil {
		panic(err)
	}

	i.col = db.Collection(role.AppName)
	return nil
}

func (i *impl) Registry(server *grpc.Server) {
	role.RegisterRPCServer(server, i)
}

func init() {
	ioc.RegistryController(&impl{})
}
