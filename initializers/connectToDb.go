package initializers

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// func ConnecttoDb() {
// 	var err error
// 	dsn := os.Getenv("DB")
// 	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatalf("Failed to connect to the database: %v", err)
// 	} else {
// 		log.Println("Successfully connected to the database")
// 	}
// }

func ConnecttoDb(){
	var err error
	dsn := "postgresql://postgres:newpass@localhost/ecomresto?sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	} else {
		log.Println("Successfully connected to the database")
	}
}
