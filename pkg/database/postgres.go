// postgres.go
package database

import (
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// // GetDB returns the database instance
// func GetDB() *gorm.DB {
//     return db
// }
func ConnectPostgres() {
    dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
    dbConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("failed to connect to db:", err)
    }
    db = dbConn
}