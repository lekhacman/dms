package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/lekhacman/dms/internal"
	"github.com/valyala/fasthttp"
)

func ToJson(appCtx *internal.AppContext) func(func(*fasthttp.RequestCtx) (interface{}, error)) func(*fasthttp.RequestCtx) {
	return func(h func(*fasthttp.RequestCtx) (interface{}, error)) func(*fasthttp.RequestCtx) {
		return func(ctx *fasthttp.RequestCtx) {
			ctx.Response.Header.Add("content-type", "application/json")
			dto, err := h(ctx)
			if err != nil {
				ctx.Error(err.Error(), 400)
				return
			}

			b, err := json.Marshal(dto)
			if err != nil {
				appCtx.Logger.Debugln("Error marshalling json")
			}

			fmt.Fprint(ctx, string(b))
		}
	}
}
