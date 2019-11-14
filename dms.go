package dms

import (
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"log"
)

func Index(ctx *fasthttp.RequestCtx) {
	_, _ = fmt.Fprint(ctx, "Document Center Service")
}

func main() {
	router := fasthttprouter.New()
	router.GET("/", Index)

	log.Fatal(fasthttp.ListenAndServe(":8001", router.Handler))
}
