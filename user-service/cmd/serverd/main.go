package main

import (
	"context"
	"database/sql"
	"log"

	"github.com/kytruong0712/goffee-shop/user-service/cmd/banner"
	"github.com/kytruong0712/goffee-shop/user-service/internal/controller/user"
	grpcSvc "github.com/kytruong0712/goffee-shop/user-service/internal/handler/grpcserver"
	"github.com/kytruong0712/goffee-shop/user-service/internal/handler/grpcserver/protogen/users"
	"github.com/kytruong0712/goffee-shop/user-service/internal/infra/config"
	"github.com/kytruong0712/goffee-shop/user-service/internal/infra/db/pg"
	grpcSvr "github.com/kytruong0712/goffee-shop/user-service/internal/infra/protocols/grpc"
	"github.com/kytruong0712/goffee-shop/user-service/internal/repository"
	"github.com/kytruong0712/goffee-shop/user-service/internal/repository/generator"

	"google.golang.org/grpc"
)

func main() {
	banner.Print()

	// Initial config
	cfg, err := initConfig()
	if err != nil {
		log.Fatal(err)
	}

	// rootCtx
	_ = context.Background()

	// Initial snowflake generator
	generator.InitSnowflakeGenerators()

	// Initial DB connection
	conn, err := pg.Connect(cfg.PGCfg.PGUrl)
	if err != nil {
		log.Fatal("[PG connection error] ", err)
	}
	defer conn.Close()

	startingGRPCServer(cfg, conn)
}

func initConfig() (config.Config, error) {
	cfg := config.NewConfig()
	if err := cfg.Validate(); err != nil {
		return config.Config{}, err
	}

	return cfg, nil
}

func initGRPCServer() *grpc.Server {
	grpcServer := grpc.NewServer()

	return grpcServer
}

func startingGRPCServer(cfg config.Config, conn *sql.DB) {
	serv := initGRPCServer()

	repo := repository.New(conn)
	userCtrl := user.New(cfg.IamConfig, repo)

	serviceServer := grpcSvc.New(userCtrl)
	users.RegisterUserServiceServer(serv, serviceServer.UserServiceServer())
	log.Printf("Started user service on %v", cfg.ServerCfg.ServerAddr)
	grpcSvr.Start(serv, cfg.ServerCfg.ServerAddr)
}
