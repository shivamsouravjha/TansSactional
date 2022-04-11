package db

import (
	"context"
	"fmt"
	"transactional/services"
	structs "transactional/struct"
	"transactional/struct/request"
)

func GetCompanyDAO(ctx context.Context, getContentRequest *request.CompanyID) (*structs.CompanyDetails, error) {
	var CompanyDetails structs.CompanyDetails
	sqlString := fmt.Sprintf("SELECT name , email from company  where companyid = \"%v\" ", getContentRequest.CompanyID)
	err := services.Dbmap.SelectOne(&CompanyDetails, sqlString)
	if err != nil {
		fmt.Println(err.Error(), sqlString)
		return nil, err
	}
	return &CompanyDetails, nil
}

func GetAllCompanyDAO(ctx context.Context, getContentRequest *request.CompanyID) (*[]structs.CompanyDetails, error) {
	var CompanyDetails []structs.CompanyDetails
	sqlString := fmt.Sprintf("SELECT name , email from company")
	_, err := services.Dbmap.Select(&CompanyDetails, sqlString)
	if err != nil {
		fmt.Println(err.Error(), sqlString)
		return nil, err
	}
	return &CompanyDetails, nil
}
