package handler

import (
	"github.com/lekhacman/dms/pkg/model"
)

func Model(_ interface{}) interface{} {
	return model.Object{}
}
