package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Init() error {
	var err error

	//initialize SQLite
	db, err = InitializeDB()
	
	if err != nil {
		return fmt.Errorf("error initialize SQLite: %v", err)
	}
	return nil
}

func GetDB() *gorm.DB {
	return db
} 

func GetLogger(p string)*Logger {
	//initialize logger
	logger := NewLogger(p)
	return logger
}