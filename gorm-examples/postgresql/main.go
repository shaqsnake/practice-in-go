package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	db, err := gorm.Open("postgres", "host=localhost user=admin dbname=taiga sslmode=disable password=magic1")
	if err != nil {
		log.Fatalf("DB connect error: %v", err)	
	}
	defer db.Close()

	db.DB().Ping()
}