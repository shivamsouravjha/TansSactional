package db

import (
	"context"
	"fmt"
	"transactional/services"
	structs "transactional/struct"
	"transactional/struct/request"
)

func GetUserDAO(ctx context.Context, getContentRequest *request.UserID) (*structs.UserDetails, error) {
	var UserDetails structs.UserDetails
	sqlString := fmt.Sprintf("SELECT user.name , user.email,company.name as companyname from usercompanyrelation join user on user.iduser =usercompanyrelation.iduser join company on usercompanyrelation.idcompany = company.companyid where usercompanyrelation.iduser=  \"%v\" ", getContentRequest.UserID)
	err := services.Dbmap.SelectOne(&UserDetails, sqlString)
	if err != nil {
		fmt.Println(err.Error(), sqlString)
		return nil, err
	}
	return &UserDetails, nil
}
