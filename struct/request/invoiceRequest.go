package request

type GetInvoice struct {
	Filter Filter `json:"filter,omitempty"`
	Sort   string `json:"sort,omitempty"`
	UserID string `json:"userId"`
	Page   int    `json:"page,omitempty"`
}

type Filter struct {
	Comperator   string `json:"comperator,omitempty"`
	Value        string `json:"value,omitempty"`
	Settled      string `json:"settled,omitempty"`
	Acknowledged string `json:"acknowledged,omitempty"`
	Position     string `json:"position,omitempty"`
}
type CreateInvoice struct {
	UserID            string   `json:"userId" binding:"required"`
	Product           []string `json:"product" binding:"required"`
	BuyingCompanyID   string   `json:"buyingcompanyid,omitempty"`
	BuyingCompanyName string   `json:"buyingcompanyname,omitempty"`
	SellingCompanyID  string   `json:"sellingcompanyid" binding:"required"`
}

type UpdateInvoce struct {
	GrandTotal string `json:"grandtotal" binding:"required"`
	InvoicesID string `json:"idinvoices" binding:"required"`
}
