package internal

import (
	"github.com/lekhacman/dms/internal/store"
	"github.com/sirupsen/logrus"
)

type AppContext struct {
	Logger *logrus.Logger
	Store  *store.Dms
}
