package db

import (
	"context"
	"fmt"
	"transactional/services"
	"transactional/struct/request"

	"golang.org/x/crypto/bcrypt"
)

func CreateCompanyDAO(ctx context.Context, getContentRequest *request.CreateCompany) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(getContentRequest.Password), 1)
	_, err := services.Dbmap.Exec("INSERT INTO company (name,email,password) VALUES(?,?,?)", getContentRequest.Name, getContentRequest.Email, hashedPassword)
	if err != nil {
		fmt.Println(err.Error())
		return "Can't create company as email is taken already"
	}
	return ""
}
