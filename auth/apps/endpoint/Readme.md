1. 功能列表注册接口

操作 ---> 接口 ---> 路由
```
r.Route(r.GET("/").To(h.QuerySecret).
    Doc("查询凭证列表").
    // 作为OpenApi的值作为展示
    Metadata(restfulspec.KeyOpenAPITags, tags).
    // 通过auth来控制是否开启认证
    Metadata("auth", true))
```

功能:
 + feature: http接口,  POST.URI   cmdb.post./cmdb/api/v1/secrets
 + service: 那个服务, cmdb
 + method: post
 + path(访问路由): /cmdb/api/v1/secrets