package config

import (
	"os"

	sqlite "github.com/glebarez/sqlite"
	"github.com/nicolas-nva/opportunities-go/schemas"
	"gorm.io/gorm"
)

//create DB and connect
func InitializeDB() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"
	//check if db exists
	_,err := os.Stat(dbPath)
	if os.IsNotExist(err){
		logger.Info("database file not found, creating...")
		//create db file and directory
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	db, err:= gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil{
		logger.ErrorF("sqlite opening error: %v", err)
		return nil, err
	}
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil{
		logger.ErrorF("sqlite automigration error: %v", err)
		return nil, err
	}
	return db, nil
}