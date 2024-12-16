package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB(connectionString string) (db *gorm.DB, err error) {
	db, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
