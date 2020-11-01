package migration

import (
	"fmt"
	"gotask/internal/app/adapter/postgresql"
	"gotask/internal/app/adapter/postgresql/model"
)

//Migrate migrates the database
func Migrate() {
	db := postgresql.Connection()
	err := db.AutoMigrate(
		model.Task{},
	)

	if err != nil {
		panic(fmt.Errorf("Fatal error in migration: %s", err))
	}
}
