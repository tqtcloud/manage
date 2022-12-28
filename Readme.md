##    :fist:后续调整
- [x] 前端 `API` 统一服务，所有的 `api` 调用写一个里面 (2022-12-26)
- [x] `error` 错误友好输出返回 (2022-12-27)
- [x] 云厂商实现一个 `rpc` 调整(2022-12-28)
- [ ] 缺少 `http` 中间件拦截器
- [ ] 整个目录结构调整

## :tada:redis 实现rpc认证

```shell
# 设置 hset key值
HSET rpc:auth:user userapi  6jKNZbEpYGeUMAifz10gOnmoty3TV
HSET rpc:auth:secret secretapi 6jKNZbEpYGeUMAifz10gOnmoty3TV
HSET rpc:auth:task taskapi 6jKNZbEpYGeUMAifz10gOnmoty3TV
HSET rpc:auth:host hostapi 6jKNZbEpYGeUMAifz10gOnmoty3TV
HSET rpc:auth:alioperator alioperatorapi 6jKNZbEpYGeUMAifz10gOnmoty3TV

# 查询
HGET rpc:auth:user userapi
# 删除
HDEL rpc:auth:host hostapi
```

### :eyes:操作

```shell
goctl model mysql ddl -src ./model/user.sql -dir ./model -c
goctl  api go -api front-api.api -dir .
goctl rpc protoc ./rpc/user.proto --go_out=./rpc/types --go-grpc_out=./rpc/types --zrpc_out=./rpc

cd $ServicePaht
goctl.exe  docker  --go user.go --port 9001 --version 1.18
```

## :blush:doc 文档生成

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

### :camera: Git操作

```shell
# Git托管删除
如果你添加.gitignore的时候，git里面已经上传了很多不需要的文件，则使用下面两个命令干掉他们
如果是文件夹：git rm -r --cached 文件夹名
如果是文件：git rm --cached 文件名
方法参考自：https://stackoverflow.com/questions/9550437/how-to-make-git-ignore-idea-files-created-by-rubymine

# Git Tag
#查看本地的Tag标签
git tag  
# 构建Tag 并输入Msg
git tag -a v0.3 -m "统一api，厂商模块" 
# 推送指定Tag到仓库
git push origin v0.3
```