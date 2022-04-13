package db

import (
	"context"
	"fmt"
	"transactional/services"
	"transactional/struct/request"

	"golang.org/x/crypto/bcrypt"
)

func UpdateCompanyDAO(ctx context.Context, updateCreatorRequest *request.CreateCompany, companyId *request.CompanyID) string {
	sql := "update `company` set "
	noupdate := true
	if len(updateCreatorRequest.Name) > 0 {
		noupdate = false
		sql += fmt.Sprintf(" name = \"%v\"", updateCreatorRequest.Name) + ","
	}
	if len(updateCreatorRequest.Email) > 0 {
		noupdate = false
		sql += fmt.Sprintf(" email = \"%v\"", updateCreatorRequest.Email) + ","
	}
	if len(updateCreatorRequest.Password) > 0 {
		noupdate = false
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(updateCreatorRequest.Password), 1)
		sql += fmt.Sprintf(" password = \"%v\"", string(hashedPassword)) + ","
	}
	sql = sql[:len(sql)-1]
	if noupdate {
		return "no data to update"
	}
	sql += fmt.Sprintf(" where companyid = %v", companyId.CompanyID)
	_, err := services.Dbmap.Exec(sql)
	if err != nil {
		fmt.Println(err.Error(), sql)
		return "unable to update company data"
	}
	return ""
}
