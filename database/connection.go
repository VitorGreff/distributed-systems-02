package database

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
		log.Fatal("Couldnt get environment variables")
	}
	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASSWORD")
	dsn := fmt.Sprintf("%v:%v@tcp(127.0.0.1:3306)/sd?charset=utf8mb4&parseTime=True&loc=Local", db_user, db_pass)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return nil, errors.New("erro ao conectar com banco")
	}

	return db, nil
}
