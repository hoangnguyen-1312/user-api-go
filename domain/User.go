package domain

import (
	"context"
	_"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uint64     `gorm:"primary_key;auto_increment" json:"id"`
	FirstName string     `gorm:"size:100;not null;" json:"first_name"`
	LastName  string     `gorm:"size:100;not null;" json:"last_name"`
	Email     string     `gorm:"size:100;not null;unique" json:"email"`
	Password  string     `gorm:"size:100;not null;" json:"password"`
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

type ProfileUser struct {
	ID        uint64 `gorm:"primary_key;auto_increment" json:"id"`
	FirstName string `gorm:"size:100;not null;" json:"first_name"`
	LastName  string `gorm:"size:100;not null;" json:"last_name"`
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *User) EncodePassword() error {
	hashPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashPassword)
	return nil
}

func (u *User) ProfileUser() interface{} {
	return &ProfileUser{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
	}
}

type UserEntity interface {
	SaveInformation(user User, ctx context.Context) (User, error)
	ShowProfile(ctx context.Context, id uint64) (User, error)
	GetUserByEmailAndPassword(user User, ctx context.Context) (User, error)
}

type UserRepository interface {
	SaveInformation(user User, ctx context.Context) (User, error)
	ShowProfile(ctx context.Context, id uint64) (User, error)
	GetUserByEmailAndPassword(user User, ctx context.Context) (User, error)
}
