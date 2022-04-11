package response

import structs "transactional/struct"

type InvoiceResponse struct {
	Message         Response
	InvoiceResponse *[]structs.InvoiceDetails
	Page            int `json:"page"`
}
