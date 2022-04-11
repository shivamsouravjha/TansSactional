package db

import (
	"context"
	"fmt"
	"transactional/services"
	"transactional/struct/request"

	"golang.org/x/crypto/bcrypt"
)

func UpdateUserDAO(ctx context.Context, updateUserRequest *request.CreateUser, userId *request.UserID) string {
	sql := "update `user` set"
	if len(updateUserRequest.Name) > 0 {
		sql += fmt.Sprintf(" name = \"%v\"", updateUserRequest.Name) + ","
	}
	if len(updateUserRequest.Email) > 0 {
		sql += fmt.Sprintf(" email = \"%v\"", updateUserRequest.Email) + ","
	}
	if len(updateUserRequest.Password) > 0 {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(updateUserRequest.Password), 1)
		sql += fmt.Sprintf(" password = \"%v\"", string(hashedPassword)) + ","
	}
	sql = sql[:len(sql)-1]
	sql += fmt.Sprintf("where iduser = %v", userId.UserID)
	_, err := services.Dbmap.Exec(sql)
	if err != nil {
		fmt.Println(err.Error(), sql)
		return "unable to update user data"
	}
	if updateUserRequest.CompanyID > 0 {
		sql = fmt.Sprintf("update `usercompanyrelation` set idcompany=\"%v\" where iduser=\"%v\"", updateUserRequest.CompanyID, userId.UserID)
		_, err = services.Dbmap.Exec(sql)
		if err != nil {
			fmt.Println(err.Error(), sql)
			return "unable to update company"
		}
	}

	return ""
}
