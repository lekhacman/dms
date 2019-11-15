package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/lekhacman/dms/pkg/model"
	"github.com/valyala/fasthttp"
	"time"
)

func Create(ctx *fasthttp.RequestCtx) {
	document := model.Document{
		uuid.New(),
		"hello world",
		"txt",
		"test",
		1,
		time.Now(),
		time.Now(),
	}

	ctx.Response.Header.Add("content-type", "application/json")

	b, err := json.Marshal(document)
	if err != nil {
		fmt.Println("error:", err)
	}
	buf := bytes.NewBuffer(b)

	fmt.Fprint(ctx, buf.String())
}
