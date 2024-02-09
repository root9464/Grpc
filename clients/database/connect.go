package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Data struct {
	DB *gorm.DB
}

var DB Data

func ConnectToDB() (Data, error) {


	dsn := "postgres://postgres.agozltsslyxdxyxewoul:8tmJqJL7exflHEFP@aws-0-eu-central-1.pooler.supabase.com:5432/postgres"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Ошибка тут")
	}
	DB = Data{
		DB: db,
	}
	return DB, nil
}