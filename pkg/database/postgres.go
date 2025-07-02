// postgres.go
package database

import (
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
    "os"
)


// // GetDB returns the database instance
// func GetDB() *gorm.DB {
//     return db
// }
func ConnectPostgres() (*gorm.DB,error){
    host := os.Getenv("DB_HOST")
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    dbname := os.Getenv("DB_NAME")
    port := os.Getenv("DB_PORT")
    log.Printf("Connecting to Postgres at %s:%s/%s with user %s", host, port, dbname, user)
    if host == "" || user == "" || password == "" || dbname == "" || port == "" {
        log.Fatal("missing one or more environment variables for database connection")
    }
    dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=disable"
    dbConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("failed to connect to db:", err)
    }
    return dbConn, nil
}