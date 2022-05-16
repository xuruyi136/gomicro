# https://github.com/CocaineCong/micro-todoList/ #Go-Micro+RabbitMQ 构建简单备忘录
# Proto 生成Go和Go-micro文件
# https://www.topgoer.com/%E5%BE%AE%E6%9C%8D%E5%8A%A1/gRPC/%E5%AE%89%E8%A3%85.html  grpc安装


cd E:\gomicro\proto\rand
protoc -I. --micro_out=. --go_out=. ./rand.proto

cd E:\gomicro\proto\user
protoc -I. --micro_out=. --go_out=. ./user.proto

ProtoBuf 版本生成错误
# 请先安装依赖 
go get -u github.com/micro/protoc-gen-micro/v2 
然后 protoc --proto_path=$GOPATH/src:. --micro_out=. --go_out=. proto/rand/rand.proto
Go Micro - https://github.com/micro/go-micro  =>https://micro.dev/getting-started
Consul - https://www.consul.io/
Gin - https://github.com/gin-gonic/gin
ProtoBuf - http://google.github.io/proto-lens/installing-protoc.html
https://github.com/protocolbuffers/protobuf/releases/tag/v3.9.0

//https://github.com/protocolbuffers/protobuf/releases

https://github.com/alibaba/sentinel-golang  限流/熔断
https://sentinelguard.io/zh-cn/docs/golang/basic-api-usage.html
https://github.com/alibaba/sentinel-golang/wiki/%E5%A6%82%E4%BD%95%E4%BD%BF%E7%94%A8

protoc-3.20.0-win64.zip
解压goroot/bin

proctoc --version
libprotoc 3.20.0


protoc-gen-go.exe  --go_out=.
protoc-gen-go-grpc.exe
protoc-gen-micro.exe  --micro_out=.

go get -d -u github.com/golang/protobuf/protoc-gen-go
go install github.com/golang/protobuf/protoc-gen-go


docker search consul
docker images -f 'reference=consul'


docker run \
-d \
-p 8500:8500 \
--name=badger \
consul agent -server -ui -node=server-1 -bootstrap-expect=1 -client=0.0.0.0


```json
{
    "username":"millssf11",
    "password":"12345",
    "email":"millyn@111.com"
}
```