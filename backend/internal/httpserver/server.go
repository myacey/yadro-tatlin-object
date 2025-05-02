package httpserver

import (
	"context"
	"log"

	"yadro-tatlin-object/internal/config"
	"yadro-tatlin-object/internal/httpserver/handler"
	"yadro-tatlin-object/internal/pkg/web"
	"yadro-tatlin-object/internal/pkg/web/middleware"
	"yadro-tatlin-object/internal/repository"
	"yadro-tatlin-object/internal/service"
	"yadro-tatlin-object/pkg/openapi"

	"github.com/gin-gonic/gin"
	oapimiddleware "github.com/oapi-codegen/gin-middleware"
)

type App struct {
	server  web.Server
	Router  *gin.Engine
	Service *service.Service
}

func New(cfg config.Config) *App {
	app := &App{
		Router: gin.Default(),
	}
	app.server = web.NewServer(cfg.HTTPServer, app.Router)

	app.initialize(cfg)

	return app
}

func (app *App) Start(ctx context.Context) error {
	return app.server.Run(ctx)
}

func (app *App) Stop(ctx context.Context) error {
	return app.server.Shutdown(ctx)
}

func (app *App) initialize(cfg config.Config) {
	repo := repository.NewPet(cfg.Repository)
	srv := service.Service{service.NewPet(repo)}
	hndler := handler.NewHandler(srv)

	app.Router.Use(middleware.RequestIDMiddleware(handler.HeaderRequestID))

	swagger, err := openapi.GetSwagger()
	if err != nil {
		log.Fatal(err)
	}

	openapi.RegisterHandlers(app.Router, hndler)
	app.Router.Use(oapimiddleware.OapiRequestValidator(swagger)) // Только на /v1/**
}
