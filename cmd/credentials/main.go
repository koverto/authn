package main

import (
	"fmt"
	"os"

	credentials "github.com/koverto/credentials/api"
	"github.com/koverto/credentials/internal/pkg/handler"
	"github.com/koverto/micro"
	"github.com/micro/go-micro/v2/config/source/env"
)

func main() {
	conf := &handler.Config{
		MongoURL: "mongodb://localhost:27017",
	}

	service, err := micro.NewService("com.koverto.svc.credentials", conf, env.NewSource(env.WithStrippedPrefix("KOVERTO")))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	h, err := handler.New(conf, service)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := credentials.RegisterCredentialsHandler(service.Server(), h); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := service.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
