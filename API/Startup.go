package API

import (
	"mediatR/Infrastructures/Database"
	"mediatR/Infrastructures/Migration"
	"mediatR/Infrastructures/Repositories"
	"os"
	"strings"
)

func Startup() {

	// connect to Database
	Database.ConnectDatabase()

	// enable migration
	enableMigration := os.Getenv("MIGRATION")
	if strings.ToUpper(enableMigration) == "TRUE" {
		Migration.MigrateTables(Database.GetDB())
	}

	//dependency injection
	Repositories.NewUnitOfWork()

	// build route config
	router := RouteConfig()
	router.Run(":8000")
}
