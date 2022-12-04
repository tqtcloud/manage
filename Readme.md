

redis 实现rpc认证
```shell
# 设置 hset key值
HSET rpc:auth:user userapi  6jKNZbEpYGeUMAifz10gOnmoty3TV
HSET rpc:auth:secret secretapi 6jKNZbEpYGeUMAifz10gOnmoty3TV
HSET rpc:auth:task taskapi 6jKNZbEpYGeUMAifz10gOnmoty3TV

 # 查询
HGET rpc:auth:user userapi
```

```shell
goctl model mysql ddl -src ./model/user.sql -dir ./model -c
goctl  api go -api user.api -dir . 
goctl rpc protoc ./rpc/user.proto --go_out=./rpc/types --go-grpc_out=./rpc/types --zrpc_out=./rpc
```

doc 文档生成
```shell
# 前往指定的服务目录
 goctl.exe  api doc --dir api/ --o api/
```

```shell
# etcd 查询
export ETCDCTL_API=3
export ENDPOINTS=192.168.0.102:2379
etcdctl --endpoints=$ENDPOINTS member list
etcdctl --endpoints=$ENDPOINTS get --prefix ""
```