package store

import (
	"database/sql"
	"fmt"
	"github.com/lekhacman/dms/pkg/model"
	"github.com/lib/pq"
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
	log *logrus.Logger
	db  *sql.DB
}

func (dms *Dms) Save(dto *model.Object) bool {
	stmt, err := dms.db.Prepare("INSERT INTO objects (id, owner_id, name, description, size, created_at, updated_at) VALUES ( $1, $2, $3, $4, $5, $6, $7 )")
	if err != nil {
		dms.log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		dto.Id.String(),
		dto.OwnerId.String(),
		dto.Name,
		dto.Description,
		dto.Size,
		pq.FormatTimestamp(dto.CreatedAt),
		pq.FormatTimestamp(dto.UpdatedAt),
	)

	if err != nil {
		dms.log.Errorf("Error inserting object %s with error %v", dto.Id, err)
		return false
	}

	return true
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
		db:  db,
		log: log,
	}

	return store
}
