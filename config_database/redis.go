package config_database
import (
	"github.com/go-redis/redis/v7"
	_"os"
	_"user-system-go/auth"
	_"fmt"
)


var Client *redis.Client


func SetupRedisDB(host, port, password string) {
	// host := os.Getenv("REDIS_HOST")
	// port := os.Getenv("REDIS_PORT")
	// password := os.Getenv("REDIS_PASSWORD")

	Client := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       0,
	})
	SetUpClient(Client)
}

func SetUpClient(client *redis.Client) {
	Client = client
}

func GetClient() *redis.Client {
	return Client 
}