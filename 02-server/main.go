package main

import (
	"context"
	"fmt"
	"net"

	"github.com/Atchett/go-grpc/02-server/echo"
	"google.golang.org/grpc"
)

// EchoServer - empty struct
type EchoServer struct{}

// Echo - function
func (e *EchoServer) Echo(ctx context.Context, req *echo.EchoRequest) (*echo.EchoResponse, error) {
	return &echo.EchoResponse{
		Response: "My Echo: " + req.Message,
	}, nil
}

func main() {
	// create a network listner
	lst, err := net.Listen("tcp", ":8000")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()

	srv := &EchoServer{}

	echo.RegisterEchoServerServer(s, srv)
	fmt.Println("Now serving at port 8000")

	err = s.Serve(lst)
	if err != nil {
		panic(err)
	}
}
