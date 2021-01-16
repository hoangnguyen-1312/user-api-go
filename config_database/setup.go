package config_database
import(
	"os"
	"github.com/joho/godotenv"
	"log"
)

func Init(){
	postgresInit()
	redisInit()
}

func postgresInit() {

	if err := godotenv.Load(); err != nil {
		log.Println("no env gotten")
	}

	host := os.Getenv("DB_HOST")
	password := os.Getenv("DB_PASSWORD")
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	SetupPostgresDB(host, password, user, dbname, port)
	// db := GetDBConnection()
}

func redisInit() {

	if err := godotenv.Load(); err != nil {
		log.Println("no env gotten")
	}

	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	password := os.Getenv("REDIS_PASSWORD")

	SetupRedisDB(host, port, password)
	// client := GetClient()
}