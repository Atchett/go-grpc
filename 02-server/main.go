package main

import (
	"context"

	"github.com/Atchett/grpc/02-server/echo"
)

// EchoServer - empty struct
type EchoServer struct{}

// Echo - function
func (e *EchoServer) Echo(context.Context, *echo.EchoRequest) (*echo.EchoResponse, error)

func main() {}
