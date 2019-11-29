package handler

import (
	"encoding/json"
	"errors"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"github.com/lekhacman/dms/internal"
	"github.com/lekhacman/dms/pkg/model"
	"github.com/valyala/fasthttp"
	"time"
)

func Create(appCtx *internal.AppContext, ctx *fasthttp.RequestCtx) (interface{}, error) {
	var dto model.Object

	err := json.Unmarshal(ctx.PostBody(), &dto)
	if err != nil {
		return nil, err
	}
	v := validator.New()
	err = v.Struct(dto)
	if err != nil {
		return nil, err
	}

	target := &model.Object{
		Id:          uuid.New(),
		OwnerId:     uuid.New(),
		Name:        dto.Name,
		Description: dto.Description,
		Size:        uint32(len([]byte(dto.Content))), // Overflow if larger than 4GB
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	ok := appCtx.Store.Save(target)
	if !ok {
		return nil, errors.New("Wa lau eh!")
	}

	return target, nil
}
