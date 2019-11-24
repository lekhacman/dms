package main

import (
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/lekhacman/dms/internal/config"
	"github.com/lekhacman/dms/internal/handler"
	log "github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"os"
)

func Index(appName string) func(ctx *fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		_, _ = fmt.Fprint(ctx, appName)
	}
}

func main() {

	conf := config.Get(fmt.Sprintf("internal/config/config.%s.toml", os.Getenv("env")))

	logMap := map[string]log.Level{
		"panic": log.PanicLevel,
		"fatal": log.FatalLevel,
		"error": log.ErrorLevel,
		"warn":  log.WarnLevel,
		"info":  log.InfoLevel,
		"debug": log.DebugLevel,
		"trace": log.TraceLevel,
	}
	logLevel, ok := logMap[conf.App.LogLevel]
	if !ok {
		log.Fatal("Log level not found")
	}
	log.SetLevel(logLevel)

	//db, err := store.NewDmsStore(store.DbSpec{
	//	"dms",
	//	os.Getenv("dbpass"),
	//	"dms",
	//	"localhost",
	//	"5432",
	//})
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//_, err = db.Query("SELECT * FROM objects")
	//if err != nil {
	//	log.Fatal(err)
	//}

	router := fasthttprouter.New()
	router.GET("/", Index(conf.Name))
	router.GET("/model", handler.Model)
	router.POST("/", handler.Create)

	log.Infof("Starting application at port %s", conf.App.Port)
	log.Fatal(fasthttp.ListenAndServe(
		fmt.Sprintf(":%s", conf.App.Port),
		router.Handler,
	))
}
