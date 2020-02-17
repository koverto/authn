package main

import (
	"fmt"
	"os"

	authn "github.com/koverto/authn/api"
	"github.com/koverto/authn/internal/pkg/handler"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/config/source/env"
)

func main() {
	service := micro.NewService(micro.Name("authn"))
	service.Init()

	conf, err := handler.NewConfig("authn", env.NewSource(env.WithStrippedPrefix("KOVERTO")))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	h, err := handler.New(conf)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := authn.RegisterAuthnHandler(service.Server(), h); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := service.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
