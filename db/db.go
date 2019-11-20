package db

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/lib/pq"

	//PostgreSQL driver
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB *gorm.DB
var DB *gorm.DB

// Open return nil
func Open() {
	var err error

	DBMS := "postgres"
	connection, err := pq.ParseURL(os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err.Error())
	}
	connection += " sslmode=" + os.Getenv("POSTGRES_SSLMODE")

	DB, err = gorm.Open(DBMS, connection)
	if err != nil {
		log.Panic(err)
	}
}

// Close return nil
func Close() {
	DB.Close()
}

// AutoMigrate return nil
func AutoMigrate() {
	if DB == nil {
		log.Panic("DB未接続")
		return
	}
	DB.AutoMigrate()
}
