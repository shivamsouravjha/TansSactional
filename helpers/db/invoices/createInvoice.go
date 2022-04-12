package db

import (
	"context"
	"fmt"
	"transactional/services"
	"transactional/struct/request"
)

func CreateInvoiceDAO(ctx context.Context, createInvoiceRequest *request.CreateInvoice) string {
	_, err := services.Dbmap.Exec("INSERT INTO invoices (user_id,selling_company_id,buying_company_id,grand_total) VALUES(?,?,?,?)", createInvoiceRequest.UserID, createInvoiceRequest.SellingCompanyID, createInvoiceRequest.BuyingCompanyID, createInvoiceRequest.GrandTotal)
	if err != nil {
		fmt.Println(err.Error())
		return "Can't create Invoice"
	}
	return ""
}
