package db

import (
	"context"
	"fmt"
	"transactional/services"
	"transactional/struct/request"

	"golang.org/x/crypto/bcrypt"
)

func CreateUserDAO(ctx context.Context, createUserRequest *request.CreateUser) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(createUserRequest.Password), 1)
	user, err := services.Dbmap.Exec("INSERT INTO user (name,email,password) VALUES(?,?,?)", createUserRequest.Name, createUserRequest.Email, hashedPassword)
	if err != nil {
		fmt.Println(err.Error())
		return "Can't create user as email is already taken"
	}
	newUserId, _ := user.LastInsertId()
	_, err = services.Dbmap.Exec("INSERT INTO usercompanyrelation (idcompany,iduser) VALUES(?,?)", createUserRequest.CompanyID, newUserId)
	if err != nil {
		fmt.Println(err.Error())
		return "No such company exsits,kindly update the detail for right company"
	}
	return ""
}
