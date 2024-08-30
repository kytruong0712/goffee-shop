package main

import (
	"context"
	"log"
	"time"

	"github.com/kytruong0712/goffee-shop/api-gateway/cmd/banner"
	userCtrl "github.com/kytruong0712/goffee-shop/api-gateway/internal/controller/user"
	"github.com/kytruong0712/goffee-shop/api-gateway/internal/gateway/user"
	"github.com/kytruong0712/goffee-shop/api-gateway/internal/infra/config"
	"github.com/kytruong0712/goffee-shop/api-gateway/internal/infra/httpserver"
	"github.com/kytruong0712/goffee-shop/api-gateway/internal/infra/iam"

	"google.golang.org/grpc"
	"google.golang.org/grpc/backoff"
)

func main() {
	banner.Print()

	// Initial config
	cfg, err := initConfig()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	ctx = iam.SetConfigToContext(ctx, cfg.IamCfg)

	userGwy, err := initUserGateway(cfg)
	if err != nil {
		log.Fatalln(err)
	}

	rtr, err := initRouter(ctx, userCtrl.New(userGwy))
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

func initUserGateway(cfg config.Config) (user.Gateway, error) {
	conn, err := grpc.NewClient(
		cfg.ServerCfg.UserServiceAddr,
		grpc.WithInsecure(),
		grpc.WithConnectParams(grpc.ConnectParams{
			Backoff: backoff.Config{
				BaseDelay:  time.Second,
				Multiplier: 1.5,
				MaxDelay:   5 * time.Second,
			},
		}))
	if err != nil {
		return nil, err
	}

	return user.New(conn), nil
}

func initRouter(ctx context.Context, userCtrl userCtrl.Controller) (router, error) {
	return router{
		ctx:         ctx,
		corsOrigins: []string{"*"},
		userCtrl:    userCtrl,
	}, nil
}
