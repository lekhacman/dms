package store

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type DbSpec struct {
	User     string
	Password string
	Name     string
	Host     string
	Port     string
}

func NewDmsStore(spec DbSpec) (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		spec.User,
		spec.Password,
		spec.Name,
		spec.Host,
		spec.Port,
		"disable",
	)

	return sql.Open("postgres", connStr)
}
