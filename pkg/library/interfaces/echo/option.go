package echo

import "github.com/99designs/gqlgen/graphql/handler"

type EchoOptions func(*EchoServer)

func WithPort(port int) EchoOptions {
	return func(e *EchoServer) {
		e.port = port
	}
}

func WithGraphQLServer(h *handler.Server) EchoOptions {
	return func(e *EchoServer) {
		e.graphQLHandler = h
	}
}
