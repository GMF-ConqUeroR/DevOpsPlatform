package impl

import (
	"auth/apps/user"
	"auth/conf"
	"auth/logger"

	"github.com/infraboard/mcube/ioc"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

var _ user.Service = (*impl)(nil)

type impl struct {
	user.UnimplementedRPCServer
	ioc.ObjectImpl
	col *mongo.Collection
}

func (i *impl) Init() error {
	db, err := conf.C().MongoDB.GetDB()
	if err != nil {
		panic(err)
	}
	i.col = db.Collection("users")
	return nil
}

func (i *impl) Name() string {
	return user.AppName
}

func (i *impl) Registry(server *grpc.Server) {
	user.RegisterRPCServer(server, i)
}

func init() {
	ioc.RegistryController(&impl{})
	logger.L().Debug().Msgf("registry %s successfully!", user.AppName)
}
