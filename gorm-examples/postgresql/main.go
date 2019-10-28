package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	db, err := gorm.Open("postgres", "host=localhost user=postgres dbname=postgres sslmode=disable password=micro")
	if err != nil {
		log.Fatalf("DB connect error: %v", err)	
	}
	defer db.Close()

	db.DB().Ping()
}