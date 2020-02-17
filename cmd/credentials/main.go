package main

import (
	"fmt"
	"os"

	credentials "github.com/koverto/credentials/api"
	"github.com/koverto/credentials/internal/pkg/handler"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/config/source/env"
)

func main() {
	service := micro.NewService(micro.Name("credentials"))
	service.Init()

	conf, err := handler.NewConfig("credentials", env.NewSource(env.WithStrippedPrefix("KOVERTO")))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	h, err := handler.New(conf)
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
