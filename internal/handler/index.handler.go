package handler

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

func Index(appName string) func(ctx *fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		_, _ = fmt.Fprint(ctx, appName)
	}
}
