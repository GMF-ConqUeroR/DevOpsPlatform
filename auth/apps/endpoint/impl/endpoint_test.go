package impl_test

import (
	"auth/apps/endpoint"
	"testing"
)

func TestRegistryEndpoints(t *testing.T) {
	in := endpoint.NewEndpointSpecSet()
	in.Add(endpoint.NewEndpointSpec("cmdb", "GET./cmdb/api/v1/secrets"))
	es, err := impl.RegistryEndpoints(ctx, in)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(es)
}
