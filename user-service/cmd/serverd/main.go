package main

import (
	"log"
	"time"

	"github.com/kytruong0712/goffee-shop/user-service/cmd/banner"
	userctrl "github.com/kytruong0712/goffee-shop/user-service/internal/controller/user"
	"github.com/kytruong0712/goffee-shop/user-service/internal/gateway/notification"
	grpcsvc "github.com/kytruong0712/goffee-shop/user-service/internal/handler/grpc"
	"github.com/kytruong0712/goffee-shop/user-service/internal/handler/grpc/protobuf"
	"github.com/kytruong0712/goffee-shop/user-service/internal/infra/config"
	"github.com/kytruong0712/goffee-shop/user-service/internal/infra/db/pg"
	grpcsvr "github.com/kytruong0712/goffee-shop/user-service/internal/infra/protocols/grpc"
	"github.com/kytruong0712/goffee-shop/user-service/internal/repository"
	"github.com/kytruong0712/goffee-shop/user-service/internal/repository/generator"

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

	// Initial snowflake generator
	generator.InitSnowflakeGenerators()

	// Initial DB connection
	conn, err := pg.Connect(cfg.PGCfg.PGUrl)
	if err != nil {
		log.Fatal("[PG connection error] ", err)
	}
	defer conn.Close()

	notificationGwy, err := initNotificationGateway(cfg)
	if err != nil {
		log.Fatalln(err)
	}

	userCtrl := userctrl.New(cfg.IamConfig, notificationGwy, repository.New(conn))

	startingGRPCServer(cfg, userCtrl)
}

func initConfig() (config.Config, error) {
	cfg := config.NewConfig()
	if err := cfg.Validate(); err != nil {
		return config.Config{}, err
	}

	return cfg, nil
}

func initNotificationGateway(cfg config.Config) (notification.Gateway, error) {
	conn, err := grpc.NewClient(
		cfg.ServerCfg.NotificationServiceAddr,
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

	return notification.New(conn), nil
}

func initGRPCServer() *grpc.Server {
	return grpc.NewServer()
}

func startingGRPCServer(cfg config.Config, userCtrl userctrl.Controller) {
	serv := initGRPCServer()
	protobuf.RegisterUserServer(serv, grpcsvc.New(userCtrl))
	log.Printf("Started user service on %v", cfg.ServerCfg.ServerAddr)
	grpcsvr.Start(serv, cfg.ServerCfg.ServerAddr)
}
