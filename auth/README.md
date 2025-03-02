项目的根路径: DevOpsPlatform(作为一个微服务工程)
需要进入auth项目内进行代码生成

1. 安装自定义proto 插件
```sh
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install github.com/favadi/protoc-go-inject-tag@latest
```

2. 代码生成
```sh
// workspace目录:  DevOpsPlatform
// 截取go.mod文件的 module 部分  module gitee.com/go-course/go11/devcloud-mini/mcenter
protoc -I=. --go_out=. --go-grpc_out=. --go_opt=module="gitee.com/go-course/go11" --go-grpc_opt=module="gitee.com/go-course/go11" auth/apps/*/pb/*.proto

protoc -I=. --go_out=. --go-grpc_out=. --go_opt=module="" --go-grpc_opt=module="" auth/apps/*/pb/*.proto
protoc-go-inject-tag -input=./auth/apps/*/model.pb.go 
```


```sh
use menter_mini
db.createUser({user: "mcenter", pwd: "123456", roles: [{ role: "dbOwner", db: "menter_mini" }]})
```

