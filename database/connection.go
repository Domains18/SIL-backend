package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
	Email string
	Password string
	Orders []Order
}


type Order struct {
	gorm.Model
	UserId uint
	Item string
	Amount float64
}




func ConnectDatabase() (*gorm.DB, error){
	dsn:= "host=localhost user=postgres password=alphauser dbname=backend port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}
	return db, nil
}

func MigrateDatabase(db *gorm.DB) error {
	err := db.AutoMigrate((&User{}), (&Order{}))
	if err != nil {
		return fmt.Errorf("failed to migrate database: %v", err)
	}
	return nil
}

