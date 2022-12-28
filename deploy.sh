#!/bin/bash

#停止服务
docker stop user
docker stop secret
docker stop host
docker stop task
docker stop front-api

#删除容器
docker rm user
docker rm secret
docker rm host
docker rm task
docker rm front-api


#删除镜像
docker rmi user:v1
docker rmi secret:v1
docker rmi task:v1
docker rmi host:v1
docker rmi front-api:v1


#删除none镜像
docker rmi $(docker images | grep "none" | awk '{print $3}')

#构建服务
docker build -t user:v1 -f service/user/rpc/Dockerfile .
docker build -t secret:v1 -f service/secret/rpc/Dockerfile .
docker build -t task:v1 -f service/task/rpc/Dockerfile .
docker build -t host:v1 -f service/provider/aliyun/host/rpc/Dockerfile .
docker build -t front-api:v1 -f service/front-api/Dockerfile .


#启动服务
docker run -itd --net=host --name=user user:v1
docker run -itd --net=host --name=secret secret:v1
docker run -itd --net=host --name=host host:v1
docker run -itd --net=host --name=task task:v1
docker run -itd --net=host --name=front-api front-api:v1