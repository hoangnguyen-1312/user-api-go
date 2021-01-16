package handler

import (
	"fmt"
	"net/http"
	_ "user-system-go/app/entity"
	"user-system-go/auth"
	"user-system-go/domain"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserEntity domain.UserEntity
	Auth       auth.AuthInterface
	Token      auth.TokenInterface
}

// func NewUserHandler(r *gin.RouterGroup, us domain.UserEntity, authJWT auth.AuthInterface, token auth.TokenInterface) {
// 	handler := &UserHandler{
// 		UserEntity: us,
// 		Token:      token,
// 		Auth:       authJWT,
// 	}
// 	r.POST("/register", handler.Register)
// 	r.POST("/login", handler.Login)
// 	r.GET("/profile", auth.AuthMiddleware(), handler.AcessProfile)
// 	r.GET("/logout", auth.AuthMiddleware(), handler.Logout)
// }

func HelloPage(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "welcome",
	})
}

func (s *UserHandler) Register(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"invalid_json": "invalid json",
		})
		return
	}
	user.EncodePassword()
	newUser, err := s.UserEntity.SaveInformation(user, c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, newUser.ProfileUser())
}

func (s *UserHandler) Login(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}

	u, userErr := s.UserEntity.GetUserByEmailAndPassword(user, c.Request.Context())
	if userErr != nil {
		c.JSON(http.StatusInternalServerError, userErr)
		return
	}
	ts, tErr := s.Token.CreateToken(u.ID)
	if tErr != nil {
		c.JSON(http.StatusUnprocessableEntity, tErr.Error())
		return
	}
	saveErr := s.Auth.CreateAuth(u.ID, ts)
	if saveErr != nil {
		c.JSON(http.StatusInternalServerError, saveErr.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": ts.AccessToken,
	})
}

func (s *UserHandler) AcessProfile(c *gin.Context) {
	metadata, _ := s.Token.ExtractTokenMetadata(c.Request)
	userId, _ := s.Auth.FetchAuth(metadata.TokenUuid)
	user, err := s.UserEntity.ShowProfile(c.Request.Context(), userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, user.ProfileUser())
}

func (s *UserHandler) Logout(c *gin.Context) {
	metadata, _ := s.Token.ExtractTokenMetadata(c.Request)
	deleteErr := s.Auth.DeleteTokens(metadata)
	if deleteErr != nil {
		c.JSON(http.StatusUnauthorized, deleteErr.Error())
		return
	}
	c.JSON(http.StatusOK, "Successfully logged out")
}
