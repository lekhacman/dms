package handler

import (
	"github.com/lekhacman/dms/internal"
	"github.com/lekhacman/dms/pkg/model"
	"github.com/valyala/fasthttp"
)

func Model(_ *internal.AppContext, ctx *fasthttp.RequestCtx) (interface{}, error) {
	return model.Object{}, nil
}
