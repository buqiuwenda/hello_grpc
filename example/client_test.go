package example

import (
	"fmt"
	"gitlab.km.com/huanzhongxi/hello_grpc/pkg/client"
	"testing"
)

func TestClient(t *testing.T) {
	cli, clean := client.NewService("127.0.0.1:8080")
	data, err := cli.Call("grpchello")
	defer clean()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(data)
	}
}