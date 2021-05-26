package Migration

import (
	"gorm.io/gorm"
	"mediatR/Domain/Entities"
)

func MigrateTables(connection *gorm.DB) error {
	return connection.AutoMigrate(
			&Entities.User{},
		)
}
