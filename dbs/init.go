package dbs

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Db  *gorm.DB
	Uri string = ""
	Err error
)

func init() {
	Db, Err = gorm.Open(postgres.Open(Uri), &gorm.Config{})

	if Err != nil {
		log.Panic(Err)
		return
	}
	Db.AutoMigrate(&ToDo{})
}
