package gonylon

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Model x
type Model struct {
	ID        uuid.UUID `gorm:"primary_key;type:uniqueidentifier;column:id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
