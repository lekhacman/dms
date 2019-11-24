package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/lekhacman/dms/pkg/model"
	"github.com/valyala/fasthttp"
)

func Model(ctx *fasthttp.RequestCtx) {
	var document model.Object

	ctx.Response.Header.Add("content-type", "application/json")
	b, err := json.Marshal(document)
	if err != nil {
		fmt.Println("error:", err)
	}
	buf := bytes.NewBuffer(b)

	fmt.Fprint(ctx, buf.String())
}
