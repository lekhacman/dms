package handler

import (
	"github.com/valyala/fasthttp"
)

func ErrorHandler(ctx *fasthttp.RequestCtx, obj interface{}) {
	ctx.Error("Internal Server Error", 500)
	ctx.Logger().Printf("reach here")
}
