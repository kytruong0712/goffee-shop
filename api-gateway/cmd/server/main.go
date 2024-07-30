package main

import (
	"context"
	"log"
	
	"github.com/kytruong0712/goffee-shop/api-gateway/internal/appconfig/httpserver"
)

func main() {
	// rootCtx
	ctx := context.Background()
	// Initial router
	rtr, err := initRouter(ctx)
	if err != nil {
		log.Fatal(err)
	}

	httpserver.Start(httpserver.Handler(
		httpserver.NewCORSConfig(rtr.corsOrigins),
		rtr.routes))
}

func initRouter(ctx context.Context) (router, error) {
	return router{
		ctx:         ctx,
		corsOrigins: []string{"*"},
	}, nil
}
