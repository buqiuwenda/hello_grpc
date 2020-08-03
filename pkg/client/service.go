package client

import (
	"context"
	pd "gitlab.km.com/huanzhongxi/hello_grpc/api/protobuf-spec/example"
	"google.golang.org/grpc"
	"log"
)


type Service interface {
	Call(name string) (msg string, err error)
}

type service struct{
	client pd.ExampleClient
}

func NewService(address string) (Service, func()){
	connect, err :=grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())

	if err !=nil{
		log.Fatalf("did not connect: %v", err)
	}

	c :=pd.NewExampleClient(connect)
	return &service{client: c},func(){
		connect.Close()
	}
}

func (s *service) Call(name string) (msg string, err error){
	r, err :=s.client.Call(context.Background(), &pd.Request{Name: name})
	if err !=nil{
		log.Fatalf("could not greet: %v", err)
		return "", err
	}

	return r.GetMsg(), nil
}