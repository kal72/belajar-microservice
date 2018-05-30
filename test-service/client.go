package main

import (
	"context"
	"fmt"

	micro "github.com/micro/go-micro"
	proto "belajar-microservice/hello-service/proto"
)


func main() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(micro.Name("test.client"))
	service.Init()

	// Create new greeter client
	greeter := proto.NewGreeterService("test", service.Client())

	// Call the greeter
	rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{Name: "Adi"})
	if err != nil {
		fmt.Println(err)
	}

	// Print response
	fmt.Println(rsp.Greeting)
}