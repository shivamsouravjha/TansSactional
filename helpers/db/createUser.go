package db

import (
	"context"
	"fmt"
	"transactional/services"
	"transactional/struct/request"

	"golang.org/x/crypto/bcrypt"
)

func CreateUserDAO(ctx context.Context, getContentRequest *request.CreateUser) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(getContentRequest.Password), 1)
	user, err := services.Dbmap.Exec("INSERT INTO user (name,email,password) VALUES(?,?,?)", getContentRequest.Name, getContentRequest.Email, hashedPassword)
	if err != nil {
		fmt.Println(err.Error())
		return "Can't create user"
	}
	newUserId, _ := user.LastInsertId()
	_, err = services.Dbmap.Exec("INSERT INTO usercompanyrelation (idcompany,iduser) VALUES(?,?)", getContentRequest.CompanyID, newUserId)
	if err != nil {
		fmt.Println(err.Error())
		return "No such company exsits"
	}
	return ""
}
