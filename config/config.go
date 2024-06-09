package config

import (
	"errors"
	"fmt"
	"os"
	"test01/internals/model"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func LoadEnvironment() {
	if err := godotenv.Load("app.env"); err != nil {
		logrus.Fatal("Error Loading .env file")
	}
}

func SetUpDatabase() (*gorm.DB, error) {
	dsn := os.Getenv(DatabaseUrl)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error connecting to the database; %v", err)
	}

	if err := RunAutoMigaration(db); err != nil {
		return nil, fmt.Errorf("error running migration")
	}

	return db, nil
}

func RunAutoMigaration(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&model.User{},
	); err != nil {
		errorMessage := fmt.Sprintf("Error migrating database %v", err)
		logrus.Error(errorMessage)
		return errors.New(errorMessage)
	}

	return nil
}
