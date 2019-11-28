package model

import (
	"github.com/google/uuid"
	"time"
)

// file size range: 0 - 4294967295 byte (4 GigaByte maximum)
type Object struct {
	Id          uuid.UUID `json:"id"`
	OwnerId     uuid.UUID `json:"owner_id"`
	Name        string    `json:"name" validate:"required"`
	Description string    `json:"description"`
	Size        uint32    `json:"size"`
	Content     string    `json:"content" validate:"required,min=1,max=4294967295"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
