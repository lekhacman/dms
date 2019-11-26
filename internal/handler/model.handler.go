package handler

import (
	"github.com/lekhacman/dms/pkg/model"
	"github.com/valyala/fasthttp"
)

func Model(ctx *fasthttp.RequestCtx) (interface{}, error) {
	return model.Object{}, nil
}
