package main

import (
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/lekhacman/dms/internal/handler"
	"github.com/valyala/fasthttp"
	"log"
)

func Index(ctx *fasthttp.RequestCtx) {
	_, _ = fmt.Fprint(ctx, "Document Center Service")
}

func main() {
	router := fasthttprouter.New()
	router.GET("/", Index)
	router.GET("/model", handler.Model)
	router.POST("/", handler.Create)

	log.Fatal(fasthttp.ListenAndServe(":8001", router.Handler))
}
