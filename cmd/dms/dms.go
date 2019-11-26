package main

import (
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/lekhacman/dms/internal/config"
	"github.com/lekhacman/dms/internal/handler"
	"github.com/lekhacman/dms/internal/service"
	"github.com/valyala/fasthttp"
	"os"
)

func Index(appName string) func(ctx *fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		_, _ = fmt.Fprint(ctx, appName)
	}
}

func NewRouter(appName string) *fasthttprouter.Router {
	router := fasthttprouter.New()
	router.GET("/", Index(appName))
	router.GET("/model", handler.Model)
	router.POST("/", handler.Create)

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

	server := NewServer(NewRouter(conf.Name), logger)

	logger.Infof("Starting application at port %s", conf.App.Port)
	logger.Fatal(server.ListenAndServe(
		fmt.Sprintf(":%s", conf.App.Port),
	))
}
