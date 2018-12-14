package entity

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	UserName string
	Password string
}

//
//func New(name string) User {
//	user := User{}
//	user.username = name
//
//	return user
//}
//
//func (user User) GetUsername() string {
//	return user.username
//}
