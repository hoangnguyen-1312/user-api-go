package router
import (
	"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/gin"
	"user-system-go/config_database"
	"user-system-go/auth"
	"log"
	"os"
	"github.com/joho/godotenv"
	_entity "user-system-go/app/entity"
	_handler "user-system-go/app/handler"
	_repo "user-system-go/app/repository"

)

func Init() {

	r := gin.Default()
	r.Use(cors.Default())
	r.Group("/v1")

	config_database.Init()
	db:= config_database.GetDBConnection()

	repoUser := _repo.NewUserRepository(db)
	entityUser := _entity.NewUserEntity(repoUser)

	redisClient := config_database.GetClient()
	authJWT := auth.NewAuth(redisClient)

	token := auth.NewToken()

	handler := &_handler.UserHandler{
		UserEntity: entityUser,
		Token:      token,
		Auth:       authJWT,
	}
	r.GET("/", _handler.HelloPage)
	r.POST("/register", handler.Register)
	r.POST("/login", handler.Login)
	r.GET("/profile", handler.AcessProfile)
	r.GET("/logout", handler.Logout)

	if err := godotenv.Load(); err != nil {
		log.Println("no env gotten")
	}

	port := os.Getenv("APP_PORT")
	r.Run(port)
}