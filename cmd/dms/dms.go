package main

import (
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/lekhacman/dms/internal/config"
	"github.com/lekhacman/dms/internal/handler"
	"github.com/valyala/fasthttp"
	"log"
	"os"
)

func Index(appName string) func(ctx *fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		_, _ = fmt.Fprint(ctx, appName)
	}
}

func main() {

	conf := config.Get(fmt.Sprintf("internal/conf/conf.%s.toml", os.Getenv("env")))

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

	log.Fatal(fasthttp.ListenAndServe(
		fmt.Sprintf(":%s", conf.App.Port),
		router.Handler,
	))
}
