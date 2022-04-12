package structs

type UserDetails struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	CompanyName string `json:"companyname"`
}
type UserID struct {
	UserID string `json:"iduser"`
}

type CompanyDetails struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type InvoiceDetails struct {
	UserID           string `json:"userId,omitempty"`
	Position         string `json:"position,omitempty"`
	GrandTotal       string `json:"grandtotal"`
	Page             string `json:"page,omitempty"`
	BuyingCompanyID  string `json:"buyingcompanyid"`
	SellingCompanyID string `json:"sellingcompanyid"`
	Acknowledged     string `json:"acknowledged"`
	Settled          string `json:"settled"`
}
