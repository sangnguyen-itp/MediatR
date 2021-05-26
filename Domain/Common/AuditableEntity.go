package Common

import (
	"time"
)

type AuditableBaseEntity struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}
