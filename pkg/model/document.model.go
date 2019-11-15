package model

import (
	"github.com/google/uuid"
	"time"
)

// file size range: 0 - 4294967295 (3.99 TB maximum)
type Document struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Ext         string    `json:"ext"`
	Description string    `json:"description"`
	Size        uint32    `json:"size"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}