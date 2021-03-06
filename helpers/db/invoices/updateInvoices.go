package db

import (
	"context"
	"fmt"
	"transactional/services"
	"transactional/struct/request"
)

func UpdateInvociesDAO(ctx context.Context, updateInvoiceRequest *request.UpdateInvoce) string {
	sql := "update `invoices` set"
	if len(updateInvoiceRequest.GrandTotal) > 0 {
		sql += fmt.Sprintf(" grand_total = \"%v\"", updateInvoiceRequest.GrandTotal) + ","
	}
	sql = sql[:len(sql)-1]
	sql += fmt.Sprintf(" where idinvoices = %v", updateInvoiceRequest.InvoicesID)
	_, err := services.Dbmap.Exec(sql)
	if err != nil {
		fmt.Println(err.Error(), sql)
		return "unable to update invoices data"
	}
	return ""
}

func SettleInvoiceDAO(ctx context.Context, updateInvoiceRequest *request.UpdateInvoce) string {
	sql := fmt.Sprintf("update `invoices` set settled = true where idinvoices=\"%v\" ", updateInvoiceRequest.InvoicesID)
	_, err := services.Dbmap.Exec(sql)
	if err != nil {
		fmt.Println(err.Error(), sql)
		return "unable to update settle invoice"
	}
	return ""
}

func AcknowledgeInvoiceDAO(ctx context.Context, updateInvoiceRequest *request.UpdateInvoce) string {
	sql := fmt.Sprintf("update `invoices` set acknowledged = true where idinvoices=\"%v\" ", updateInvoiceRequest.InvoicesID)
	_, err := services.Dbmap.Exec(sql)
	if err != nil {
		fmt.Println(err.Error(), sql)
		return "unable to update acknowledge invoice"
	}
	return ""
}
