package config_database

import (
	"fmt"
	"user-system-go/domain"
	_"user-system-go/app/repository"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_"os"
	"github.com/joho/godotenv"
	"log"

)

var db *gorm.DB
var port string

type Repositories struct {
	User domain.UserRepository
	db   *gorm.DB
}

func SetupPostgresDB(host, password, user, dbname, port string){
	if err := godotenv.Load(); err != nil {
		log.Println("no env gotten")
	}

	// host := os.Getenv("DB_HOST")
	// password := os.Getenv("DB_PASSWORD")
	// user := os.Getenv("DB_USER")
	// dbname := os.Getenv("DB_NAME")
	// port := os.Getenv("DB_PORT")

	connectionParams :=  fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, password)
	
	db, err := gorm.Open("postgres", connectionParams)
	if err != nil {
		fmt.Println("Failed to connect to Postgres database!", err)
		return 
	}
	if !db.HasTable(&domain.User{}) {
		db.CreateTable(&domain.User{})
	}

	if err != nil {
		fmt.Println("Failed to connect to database!", err)
		return 
	}
	fmt.Println("Connect successfully to database!")

	db.AutoMigrate(&domain.User{})
	SetUpDBConnection(db)
}

func SetUpDBConnection(DB *gorm.DB) {
	db = DB
}

func GetDBConnection() *gorm.DB {
	return db
}
