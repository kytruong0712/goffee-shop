package main

import (
	"log"

	"github.com/kytruong0712/goffee-shop/notification-service/cmd/banner"
	"github.com/kytruong0712/goffee-shop/notification-service/internal/controller/otp"
	"github.com/kytruong0712/goffee-shop/notification-service/internal/gateway/twilio"
	grpcsvc "github.com/kytruong0712/goffee-shop/notification-service/internal/handler/grpc"
	"github.com/kytruong0712/goffee-shop/notification-service/internal/handler/grpc/protobuf"
	"github.com/kytruong0712/goffee-shop/notification-service/internal/infra/config"
	"github.com/kytruong0712/goffee-shop/notification-service/internal/infra/httpserver"
	grpcsvr "github.com/kytruong0712/goffee-shop/notification-service/internal/infra/protocols/grpc"

	"google.golang.org/grpc"
)

func main() {
	banner.Print()

	// Initial config
	cfg, err := initConfig()
	if err != nil {
		log.Fatal(err)
	}

	twilioGwy := initTwilioGateway(cfg.ServerCfg)
	optCtrl := otp.New(twilioGwy)

	startingGRPCServer(cfg, optCtrl)
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

func startingGRPCServer(cfg config.Config, otpCtrl otp.Controller) {
	serv := initGRPCServer()
	protobuf.RegisterNotificationServer(serv, grpcsvc.New(otpCtrl))
	log.Printf("Started otp service on %v", cfg.ServerCfg.ServerAddr)
	grpcsvr.Start(serv, cfg.ServerCfg.ServerAddr)
}

func initTwilioGateway(cfg httpserver.Config) twilio.Gateway {
	return twilio.New(cfg)
}
