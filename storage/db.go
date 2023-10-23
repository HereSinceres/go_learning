package storage

import (
	"log"

	"github.com/xukai-echo/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func NewDB(params ...string) *gorm.DB {

	var err error
	conString := config.GetMySQLConnectionString()
	log.Print(conString)

	DB, err = gorm.Open(mysql.Open(conString))

	if err != nil {
		log.Panic(err)
	}
	return DB

}

func GetDBInstance() *gorm.DB {
	return DB
}
