package models

import (
	"crypto/rand"
	"fmt"
	"golang.org/x/crypto/scrypt"
	"time"
)

type User struct {
	ID       string
	Name     string
	Password string
	Salt     string
	RegDate  time.Time
}

func salt() (string, error) {
	salt := make([]byte, 64)
	/*
		// 这个位置调用 io.ReadFull 多余，rand.Read 底层就是调用了 io.ReadFull
		if _, err := io.ReadFull(rand.Reader, salt); err != nil {
			return "", err
		}
	*/
	if _, err := rand.Reader.Read(salt); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", salt), nil
}

func hash(pwd string, salt string) (string, error) {
	hash, err := scrypt.Key([]byte(pwd), []byte(salt), 16384, 8, 1, 64)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hash), nil
}

// 新建用户结构体
func NewUser(form *RegisterForm, t time.Time) (*User, error) {
	salt, err := salt()
	if err != nil {
		return nil, err
	}
	hash, err := hash(form.Password, salt)
	if err != nil {
		return nil, err
	}
	user := User{
		ID:       form.Phone,
		Name:     form.Name,
		Password: hash,
		Salt:     string(salt),
		RegDate:  t,
	}
	return &user, nil
}
