package main

import (
	"context"
	"github.com/kytruong0712/goffee-shop/api-gateway/internal/controller/user"
	"github.com/kytruong0712/goffee-shop/api-gateway/internal/gateway/grpcclient"
	"google.golang.org/grpc"
	"log"

	"github.com/kytruong0712/goffee-shop/api-gateway/cmd/banner"
	"github.com/kytruong0712/goffee-shop/api-gateway/internal/infra/config"
	"github.com/kytruong0712/goffee-shop/api-gateway/internal/infra/httpserver"
)

func main() {
	banner.Print()
	// rootCtx
	ctx := context.Background()

	// Initial config
	cfg, err := initConfig()
	if err != nil {
		log.Fatal(err)
	}

	client, err := initGRPCClient(cfg)
	if err != nil {
		log.Fatalln(err)
	}

	rtr, err := initRouter(ctx, user.New(client))
	if err != nil {
		log.Fatal(err)
	}

	httpserver.Start(httpserver.Handler(httpserver.NewCORSConfig(rtr.corsOrigins), rtr.routes),
		cfg.ServerCfg)
}

func initConfig() (config.Config, error) {
	cfg := config.NewConfig()
	if err := cfg.Validate(); err != nil {
		return config.Config{}, err
	}

	return cfg, nil
}

func initGRPCClient(cfg config.Config) (grpcclient.ServiceClient, error) {
	conn, err := grpc.NewClient(cfg.ServerCfg.UserServiceAddr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return grpcclient.New(conn), nil
}

func initRouter(ctx context.Context, userCtrl user.Controller) (router, error) {
	return router{
		ctx:         ctx,
		corsOrigins: []string{"*"},
		userCtrl:    userCtrl,
	}, nil
}
