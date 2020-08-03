# go-grpc练习

go-grpc

 ### 安装proto (mac) 
 
   命令：`brew install protoc`
   ### 安装 goalng 的proto编译支持
    `go get -u github.com/golang/protobuf/protoc-gen-go`
   ### 安装 gRPC包
    `go get -u google.golang.org/grpc` 
    
   ### 安装 cobra 
   1. 设置代理：
     ` export GOPROXY=https://goproxy.cn`
     ` export GO111MODULE=on`
   
    `go get -u -v github.com/spf13/cobra/cobra`
    
### 生成grpc文件
  命令：`cobra add grpc`    