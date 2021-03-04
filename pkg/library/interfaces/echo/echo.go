package echo

import (
	"fmt"

	"github.com/graphql-go/handler"
	"github.com/labstack/echo"
)

type EchoServer struct {
	e              *echo.Echo
	port           int
	graphQLHandler *handler.Handler
}

func NewEcho(opts ...EchoOptions) (*EchoServer, error) {
	e := echo.New()

	echoServer := &EchoServer{
		e: e,
	}

	for _, opt := range opts {
		opt(echoServer)
	}

	e.GET("/graphql", echo.WrapHandler(echoServer.graphQLHandler))
	e.POST("/graphql", echo.WrapHandler(echoServer.graphQLHandler))

	return echoServer, nil
}

func (e *EchoServer) Start() error {
	return e.e.Start(fmt.Sprintf(`:%d`, e.port))
}
