package service

import (
	"context"
	pd "gitlab.km.com/huanzhongxi/hello_grpc/api/protobuf-spec/example"
	"log"
)

func NewServer() pd.ExampleServer{
	return &server{}
}

type server struct{}


func (e *server) Call(ctx context.Context, req *pd.Request) (rsp *pd.Response, err error){
	log.Println("Received Comments.Call request")

	rsp = &pd.Response{}
	rsp.Msg ="Hello" +req.Name
	return rsp, nil
}

func (e *server) Stream(req *pd.StreamingRequest, stream pd.Example_StreamServer) error{
	log.Printf("Received Comments.Stream request with count: %d\n", req.Count)

	for i:=0; i<int(req.Count); i++{
		log.Printf("Responding: %d\n", i)
		if err :=stream.Send(&pd.StreamingResponse{
			Count: int64(i),
		}); err!=nil{
			return err
		}
	}

	return nil
}

func (e *server) PingPong(stream pd.Example_PingPongServer) error {
	for{
		req, err :=stream.Recv();
		if err!=nil{
			return err
		}
		log.Printf("Got ping %v\n", req.Stroke)
		if err :=stream.Send(&pd.Pong{Stroke: req.Stroke}); err!=nil{
			return err
		}
	}
}