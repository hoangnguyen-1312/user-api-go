package main
import (
	"user-system-go/router"
	"user-system-go/config_database"
)
func main() {

	config_database.Init()
	router.Init()

}