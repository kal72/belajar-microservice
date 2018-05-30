package main

import (
	"log"
	"github.com/micro/go-micro"
	proto "belajar-microservice/hello-service/proto"
	"golang.org/x/net/context"
)
type Test struct{}
func (g *Test) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.Greeting = "Hello test" + req.Name
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("test"),
		micro.Version("latest"),
	)
	service.Init()
	proto.RegisterGreeterHandler(service.Server(), new(Test))
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
