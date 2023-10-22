package config

import (
	"chatapp/model"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connect(migrate bool) (*gorm.DB, error) {
	dns := fmt.Sprintf(
		`	host=%s
			user=%s
			password=%s
			dbname=%s
			port=%s
			sslmode=disable`,
		dbHost,
		dbUser,
		dbPassword,
		dbName,
		dbPort,
	)

	var err error

	db, err = gorm.Open(postgres.New(postgres.Config{
		DSN: dns,
	}), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	if migrate {
		log.Println("Start migrate")
		err := db.AutoMigrate(
			&model.Credential{},
			&model.Profile{},
			&model.Role{},
			&model.TemporaryCredential{},
			&model.Text{},
			&model.Group{},
			&model.GroupProfile{},
			&model.Message{},
			&model.ChatSection{},
		)

		if err != nil {
			return nil, err
		}

		log.Println("Migrate done!")
	}

	return db, nil
}
