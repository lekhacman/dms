package middleware

import "github.com/valyala/fasthttp"

type Middleware func(handler fasthttp.RequestHandler) fasthttp.RequestHandler

func Adapt(h fasthttp.RequestHandler, adapters ...Middleware) fasthttp.RequestHandler {
	for _, adapter := range adapters {
		h = adapter(h)
	}

	return h
}
