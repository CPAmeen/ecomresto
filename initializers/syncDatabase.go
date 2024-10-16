// package initializers

// import (
// 	"ecomresto/models"
// 	"fmt"
// )

//	func SyncDatabase() {
//		err := DB.AutoMigrate(&models.Users{})
//		if err != nil {
//			fmt.Println("Migration failed:", err)
//		} else {
//			fmt.Println("Migration successful")
//		}
//	}
//
// initializers/syncDatabase.go

package initializers

import (
	"ecomresto/models"

	"gorm.io/gorm"
)

func SyncDatabase(db *gorm.DB) {
	models.MigrateUsers(db)
}
