package config

import (
	"log"
	"server/internal/auth/models/entity"
)

func RunMigrations() {
	database := GetDB()

	//Arrange all table structures in one interface for maintainance
	tables := []interface{}{
		entity.User{},
	}

	for _, dbtables := range tables {
		if err := database.AutoMigrate(dbtables); err != nil {
			log.Fatalf("[Corethreads]  Failed to Migrate tables %T: %v", tables, err)
		}
	}
}
