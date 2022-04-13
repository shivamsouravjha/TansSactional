package db

import (
	"context"
	"fmt"
	"transactional/services"
	"transactional/struct/request"
)

func CreateProductDAO(ctx context.Context, createProductRequest *request.CreateProduct, userID string) string {
	_, err := services.Dbmap.Exec("INSERT INTO product (id_company,productname,productprice) select idcompany,?,? from usercompanyrelation where iduser = ?", createProductRequest.ProductName, createProductRequest.Price, userID)
	if err != nil {
		fmt.Println(err.Error())
		return "Can't create Company"
	}
	return ""
}
