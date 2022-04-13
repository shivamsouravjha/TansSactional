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
	var products []string
	sqlString := fmt.Sprintf("SELECT productname from product where idproduct in  %v", arrayString)
	_, err := services.Dbmap.Select(&products, sqlString)
	if err != nil {
		fmt.Println(err.Error(), sqlString)
		return "Can't create Invoice"
	}
	sqlString = fmt.Sprintf("INSERT INTO invoices (user_id,selling_company_id,buying_company_id,products,grand_total) SELECT \"%v\",\"%v\",\"%v\",\"%v\",sum(productprice) from product where idproduct in %v", createInvoiceRequest.UserID, createInvoiceRequest.SellingCompanyID, createInvoiceRequest.BuyingCompanyID, products, arrayString)
	_, err = services.Dbmap.Exec(sqlString)
	if err != nil {
		fmt.Println(err.Error(), sqlString)
		return "Can't create Invoice"
	}
	return ""
}
