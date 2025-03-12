package impl

import (
	"auth/apps/endpoint"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func (i *impl) RegistryEndpoints(ctx context.Context, in *endpoint.EndpointSpecSet) (*endpoint.EndpointSet, error) {
	// 添加新的service endpoint前，先清除所有旧的
	_, err := i.col.DeleteMany(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	endpointSet := endpoint.NewEndpointSet(in)
	_, err = i.col.InsertMany(ctx, endpointSet.ToMongoDoc())
	if err != nil {
		return nil, err
	}
	return endpointSet, nil
}
