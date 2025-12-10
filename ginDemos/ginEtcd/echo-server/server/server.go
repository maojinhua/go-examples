package server

import (
	"context"
	"fmt"

	"github.com/maojinhua/ginEtcd/echo"
)


type EchoService struct{
	echo.UnimplementedEchoServer
}

func (EchoService) UnaryEcho(ctx context.Context,in  *echo.EchoMessage) (*echo.EchoMessage, error) {
	fmt.Print("server recv ",in.Message)
	return &echo.EchoMessage{
		Message: "server send",
	}, nil
}