package Database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func ConnectDatabase() {

	driver := os.Getenv("DB_DRIVER")
	database := os.Getenv("DB_NAME")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	if driver == "postgres" {

		url := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
			username,
			password,
			host,
			port,
			database)

		db, err := gorm.Open(postgres.Open(url), &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
		})

		if err != nil {
			panic(err)
		}

		DB = db.Session(&gorm.Session{SkipDefaultTransaction: true})
	}
}

// GetDB helps you to get a connection
func GetDB() *gorm.DB {
	return DB
}
