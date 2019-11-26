package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
)

func AsJson(logger *logrus.Logger, h func(body interface{}) (res interface{})) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {

		ctx.Response.Header.Add("content-type", "application/json")
		dto := h(nil)
		b, err := json.Marshal(dto)
		if err != nil {
			logger.Debugln("Error marshalling json")
		}
		buf := bytes.NewBuffer(b)

		fmt.Fprint(ctx, buf.String())
	}
}
