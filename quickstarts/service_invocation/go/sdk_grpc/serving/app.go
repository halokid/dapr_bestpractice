package main

import (
	"context"
	"errors"
	"log"

	"github.com/dapr/go-sdk/service/common"
	daprd "github.com/dapr/go-sdk/service/grpc"
)

// TODO: dapr run --app-id serving --app-protocol grpc --app-port 8080 --dapr-grpc-port 3500 --log-level debug -- go run .
func main() {
	// create a Dapr service (e.g. ":8080", "0.0.0.0:8080", "10.1.1.1:8080" )
	s, err := daprd.NewService(":8080")
	if err != nil {
		log.Fatalf("failed to start the server: %v", err)
	}

	// add a service to service invocation handler
	if err := s.AddServiceInvocationHandler("echo", echoHandler); err != nil {
		log.Fatalf("error adding invocation handler: %v", err)
	}

	// start the server
	if err := s.Start(); err != nil {
		log.Fatalf("server error: %v", err)
	}
}

func echoHandler(ctx context.Context, in *common.InvocationEvent) (out *common.Content, err error) {
	if in == nil {
		err = errors.New("nil invocation parameter")
		return
	}
	log.Printf(
		"echo - ContentType:%s, Verb:%s, QueryString:%s, %s",
		in.ContentType, in.Verb, in.QueryString, in.Data,
	)
	out = &common.Content{
		Data:        in.Data,
		ContentType: in.ContentType,
		DataTypeURL: in.DataTypeURL,
	}
	return
}


