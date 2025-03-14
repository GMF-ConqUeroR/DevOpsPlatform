package endpoint

import (
	"fmt"

	"github.com/infraboard/mcube/pb/resource"
)

const AppName = "endpoints"

// Role需要应用Endpoint的Id, 要保持ID不变，Endpint的Id不能随机生成
func (s *EndpointSpec) GenId() string {
	return fmt.Sprintf("%s.%s", s.Service, s.Path)
}

func NewEndpoint(spec *EndpointSpec) *Endpoint {
	meta := resource.NewMeta()
	meta.Id = spec.GenId()
	return &Endpoint{
		Meta: meta,
		Spec: spec,
	}
}

func (es *EndpointSet) Add(items ...*Endpoint) {
	es.Items = append(es.Items, items...)
}

func NewEndpointSet(specs *EndpointSpecSet) *EndpointSet {
	set := &EndpointSet{
		Items: []*Endpoint{},
	}
	for i := range specs.Items {
		spec := specs.Items[i]
		set.Add(NewEndpoint(spec))
	}
	return set
}

func (es *EndpointSet) ToMongoDoc() (docs []interface{}) {
	for e := range es.Items {
		docs = append(docs, es.Items[e])
	}
	return
}

func NewEndpointSpec(svc, path string) *EndpointSpec {
	return &EndpointSpec{
		Service: svc,
		Path:    path,
	}
}

func NewEndpointSpecSet() *EndpointSpecSet {
	return &EndpointSpecSet{
		Items: []*EndpointSpec{},
	}
}

func (es *EndpointSpecSet) Add(items ...*EndpointSpec) {
	es.Items = append(es.Items, items...)
}

func GetValueFromMeta(meta map[string]interface{}, key string) string {
	v := meta[key]
	if v == nil {
		return ""
	}

	return v.(string)
}
