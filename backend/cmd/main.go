package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"yadro-tatlin-object/internal/config"
	"yadro-tatlin-object/internal/httpserver"
)

var cfgPath = flag.String("f", "configs/config.yaml", "path to the go's config")

func main() {
	flag.Parse()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	cfg, err := config.LoadConfig(*cfgPath)
	if err != nil {
		log.Fatal(err)
	}

	app := httpserver.New(cfg)

	go func() {
		<-ctx.Done()

		app.Stop(ctx)
	}()

	if err := app.Start(ctx); err != nil {
		log.Printf("http server returned error: %v", err)
	}
}
