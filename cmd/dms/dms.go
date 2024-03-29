package main

import (
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/lekhacman/dms/internal"
	"github.com/lekhacman/dms/internal/config"
	"github.com/lekhacman/dms/internal/handler"
	"github.com/lekhacman/dms/internal/service"
	"github.com/lekhacman/dms/internal/store"
	"github.com/lekhacman/dms/pkg/middleware"
	"github.com/valyala/fasthttp"
	"os"
)

func NewRouter(appCtx *internal.AppContext, appName string) *fasthttprouter.Router {
	toJsonMid := middleware.Json(appCtx)
	router := fasthttprouter.New()
	router.GET("/", handler.Index(appName))
	router.GET("/model", toJsonMid(handler.Model))
	router.POST("/", toJsonMid(handler.Create))

	return router
}

func NewServer(router *fasthttprouter.Router, logger fasthttp.Logger) *fasthttp.Server {
	server := &fasthttp.Server{
		MaxRequestBodySize: 10 * 1024 * 1024,
		Handler:            router.Handler,
		Logger:             logger,
	}

	return server
}

func main() {

	conf := config.Get(fmt.Sprintf("internal/config/config.%s.toml", os.Getenv("env")))

	logger := service.NewLogger(conf.App.LogLevel)
	s := store.New(logger, conf.Db)
	appCtx := &internal.AppContext{
		Logger: logger,
		Store:  s,
	}

	server := NewServer(NewRouter(appCtx, conf.Name), logger)

	logger.Infof("Starting application at port %s", conf.App.Port)
	logger.Fatal(server.ListenAndServe(
		fmt.Sprintf(":%s", conf.App.Port),
	))
}
