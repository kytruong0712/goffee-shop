package main

import (
	"log"

	"github.com/kytruong0712/goffee-shop/menu-service/cmd/banner"
	categoryctrl "github.com/kytruong0712/goffee-shop/menu-service/internal/controller/category"
	grpcsvc "github.com/kytruong0712/goffee-shop/menu-service/internal/handler/grpc"
	"github.com/kytruong0712/goffee-shop/menu-service/internal/handler/grpc/protobuf"
	"github.com/kytruong0712/goffee-shop/menu-service/internal/infra/config"
	"github.com/kytruong0712/goffee-shop/menu-service/internal/infra/db/pg"
	grpcsvr "github.com/kytruong0712/goffee-shop/menu-service/internal/infra/protocols/grpc"
	"github.com/kytruong0712/goffee-shop/menu-service/internal/repository"
	"github.com/kytruong0712/goffee-shop/menu-service/internal/repository/generator"

	"google.golang.org/grpc"
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

	repo := repository.New(conn)
	categoryCtrl := categoryctrl.New(repo)
	startingGRPCServer(cfg, categoryCtrl)
}

func initConfig() (config.Config, error) {
	cfg := config.NewConfig()
	if err := cfg.Validate(); err != nil {
		return config.Config{}, err
	}

	return cfg, nil
}

func initGRPCServer() *grpc.Server {
	return grpc.NewServer()
}

func startingGRPCServer(cfg config.Config, categoryCtrl categoryctrl.Controller) {
	serv := initGRPCServer()
	protobuf.RegisterMenuServer(serv, grpcsvc.New(categoryCtrl))
	log.Printf("Started menu service on %v", cfg.ServerCfg.ServerAddr)
	grpcsvr.Start(serv, cfg.ServerCfg.ServerAddr)
}
