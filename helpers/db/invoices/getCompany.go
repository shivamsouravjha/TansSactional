package db

import (
	"context"
	"fmt"
	"strconv"
	"transactional/services"
	structs "transactional/struct"
	"transactional/struct/request"
)

func GetInvoicesDAO(ctx context.Context, getInvoice *request.GetInvoice) (*[]structs.InvoiceDetails, string) {
	// for buyer position 0,for seller 1

	var InvoiceDetails []structs.InvoiceDetails
	sqlString := fmt.Sprintf(`SELECT grand_total as grandtotal,buying_company_id as buyingcompanyid,selling_company_id as 
	sellingcompanyid,acknowledged,settled, if(usercompanyrelation.idcompany=invoices.selling_company_id,"seller","buyer") as position from invoices 
	join usercompanyrelation on invoices.user_id = usercompanyrelation.iduser 
	join company on usercompanyrelation.idcompany = company.companyid where usercompanyrelation.iduser= %v `, getInvoice.UserID)
	if len(getInvoice.Filter.Comperator) != 0 && len(getInvoice.Filter.Value) != 0 {
		if getInvoice.Filter.Comperator == "greater" {
			sqlString += fmt.Sprintf("&& grand_total > \"%v\" ", getInvoice.Filter.Value)
		} else if getInvoice.Filter.Comperator == "less" {
			sqlString += fmt.Sprintf("&& grand_total < \"%v\" ", getInvoice.Filter.Value)
		} else if getInvoice.Filter.Comperator == "equal" {
			sqlString += fmt.Sprintf("&& grand_total = \"%v\" ", getInvoice.Filter.Value)
		}
	}
	settled, err := strconv.ParseBool(getInvoice.Filter.Settled)
	if err == nil {
		sqlString += fmt.Sprintf("&& settled = %v ", settled)
	}
	position, err := strconv.ParseBool(getInvoice.Filter.Position)
	if err == nil {
		if position {
			sqlString += fmt.Sprintf("&& usercompanyrelation.idcompany=invoices.selling_company_id ")
		} else {
			sqlString += fmt.Sprintf("&& usercompanyrelation.idcompany!=invoices.selling_company_id ")
		}
	}
	acknowledged, err := strconv.ParseBool(getInvoice.Filter.Acknowledged)
	if err == nil {
		sqlString += fmt.Sprintf("&& acknowledged = %v ", acknowledged)
	}

	if len(getInvoice.Sort) != 0 {
		sqlString += fmt.Sprintln("ORDER BY grand_total ", getInvoice.Sort)
	}
	if getInvoice.Page == 0 {
		sqlString += fmt.Sprintln("limit 5 offset 0 ")
	} else {
		sqlString += fmt.Sprintln("limit 5 offset ", getInvoice.Page*5)
	}
	_, err = services.Dbmap.Select(&InvoiceDetails, sqlString)
	if err != nil || len(InvoiceDetails) == 0 {
		fmt.Println("No such invoice exists", sqlString)
		return nil, "No such invoice exists"
	}
	for index, key := range InvoiceDetails {
		if key.Settled == "0" {
			InvoiceDetails[index].Settled = "Unsettled"
		} else {
			InvoiceDetails[index].Settled = "Settled"
		}
		if key.Acknowledged == "0" {
			InvoiceDetails[index].Acknowledged = "Un-acknowledged"
		} else {
			InvoiceDetails[index].Acknowledged = "Acknowledged"
		}
	}
	return &InvoiceDetails, ""
}
