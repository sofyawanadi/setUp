// postgres.go
package database

import (
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


// // GetDB returns the database instance
// func GetDB() *gorm.DB {
//     return db
// }
func ConnectPostgres() (*gorm.DB,error){
    dsn := "host=localhost user=postgres password=postgres dbname=set_up port=5432 sslmode=disable"
    dbConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("failed to connect to db:", err)
    }
    return dbConn, nil
}