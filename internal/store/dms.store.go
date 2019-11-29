package store

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type DbSpec struct {
	User     string
	Password string
	Name     string
	Host     string
	Port     string
}

type Dms struct {
	db *sql.DB
}

func New(log *logrus.Logger, spec DbSpec) *Dms {
	connStr := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		spec.User,
		spec.Password,
		spec.Name,
		spec.Host,
		spec.Port,
		"disable",
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Database connection fail")
	}

	store := &Dms{
		db: db,
	}

	return store
}
