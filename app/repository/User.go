package repository

import (
	"context"
	"github.com/jinzhu/gorm"
	"user-system-go/domain"
	"strings"
	"golang.org/x/crypto/bcrypt"
	"errors"
)

type UserRepository struct {
	Conn *gorm.DB
}

// NewMysqlArticleRepository will create an object that represent the article.Repository interface
func NewUserRepository(Conn *gorm.DB) domain.UserRepository {
	return &UserRepository{Conn}
}

func (r *UserRepository) SaveInformation(user domain.User, ctx context.Context) (domain.User, error) {
	err := r.Conn.Debug().Create(&user).Error
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "Duplicate") {
			return domain.User{}, err
		}
		return domain.User{}, err
	}
	return user, nil
}

func (r *UserRepository) ShowProfile(ctx context.Context, id uint64) (domain.User, error) {
	var user domain.User
	err := r.Conn.Debug().Where("id = ?", id).Take(&user).Error
	if err != nil {
		return domain.User{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return domain.User{}, errors.New("user not found")
	}
	return user, nil
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (r *UserRepository) GetUserByEmailAndPassword(u domain.User, ctx context.Context) (domain.User, error) {
	var user domain.User
	err := r.Conn.Debug().Where("email = ?", u.Email).Take(&user).Error
	if gorm.IsRecordNotFoundError(err) {
		return domain.User{}, err
	}
	if err != nil {
		return domain.User{}, err
	}
	err = VerifyPassword(user.Password, u.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return domain.User{}, err
	}
	return user, nil
}