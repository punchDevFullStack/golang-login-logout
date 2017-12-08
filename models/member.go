package models

import (
	"fmt"

	"github.com/thanamat/golang-login-logout/databases/mysql"
)

type Member struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `gorm:"primary_key" json:"email"`
	Password  string `json:"password"`
}

func CheckMemberLogin(email, password string) (*Member, error) {
	db := mysql.Connect()
	var member Member
	err := db.Where("Email = ? AND Password = ?", email, password).First(&member).Error
	fmt.Println(member, err)
	if err != nil {
		return nil, err
	}
	return &member, nil
}
