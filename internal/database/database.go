package database

import (
	"drone-backend/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Initialize() error {
    dsn := "host=localhost user=postgres password=123456 dbname=dronedb port=5432 sslmode=disable"
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return err
    }

    // Drop and recreate the tables to ensure proper schema
    // err = DB.Migrator().DropTable(&models.Drone{}, &models.Attribute{}, &models.Policy{})
    // if err != nil {
    //     return err
    // }

    err = DB.AutoMigrate(
        &models.Drone{},
        &models.Attribute{},
        &models.Policy{},
        &models.History{},
    )
    if err != nil {
        return err
    }

    return nil
}
