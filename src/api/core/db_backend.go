package core

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func (db DBBackend) CreateTables() error{
	if !db.Db.HasTable(&Location{}){
		db.Db.AutoMigrate(&Location{})
	}

	if !db.Db.HasTable(&Driver{}){
		db.Db.AutoMigrate(&Driver{})
	}
	return nil
}

func NewDBBackend() (*DBBackend, error) {
	dbUsername := "root"
	dbPassword := "root"
	dbName := "locations"
	dbHost := "localhost:3306"
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", dbUsername, dbPassword, dbHost, dbName)

	db, err	:= gorm.Open("mysql", connectionString)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return &DBBackend{Db: db}, nil
}

type DBBackend struct {
	Db *gorm.DB
}
