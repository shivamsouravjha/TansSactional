package db

import (
	"context"
	"fmt"
	"transactional/services"
	"transactional/struct/request"
)

func CreateInvoiceDAO(ctx context.Context, createInvoiceRequest *request.CreateInvoice) string {
	arrayString := "("
	for _, key := range createInvoiceRequest.Product {
		arrayString += key + ","
	}
	arrayString = arrayString[:len(arrayString)-1]
	arrayString += ")"
	sqlString := fmt.Sprintf("INSERT INTO invoices (user_id,selling_company_id,buying_company_id,grand_total) SELECT \"%v\",\"%v\",\"%v\",sum(productprice) from product where idproduct in %v", createInvoiceRequest.UserID, createInvoiceRequest.SellingCompanyID, createInvoiceRequest.BuyingCompanyID, arrayString)
	_, err := services.Dbmap.Exec(sqlString)
	if err != nil {
		fmt.Println(err.Error())
		return "Can't create Invoice"
	}
	return ""
}
