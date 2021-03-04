package echo

import "github.com/graphql-go/handler"

type EchoOptions func(*EchoServer)

func WithPort(port int) EchoOptions {
	return func(e *EchoServer) {
		e.port = port
	}
}

func WithGraphQLServer(h *handler.Handler) EchoOptions {
	return func(e *EchoServer) {
		e.graphQLHandler = h
	}
}
