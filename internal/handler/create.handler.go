package handler

import (
	"github.com/google/uuid"
	"github.com/lekhacman/dms/pkg/model"
	"github.com/valyala/fasthttp"
	"time"
)

func Create(ctx *fasthttp.RequestCtx) (interface{}, error) {
	return model.Object{
		uuid.New(),
		uuid.New(),
		"hello world",
		"test",
		1,
		"",
		time.Now(),
		time.Now(),
	}, nil
}
