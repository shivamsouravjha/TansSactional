package request

type GetInvoice struct {
	Filter Filter `json:"filter,omitempty"`
	Sort   string `json:"sort,omitempty"`
	UserID string `json:"userId"`
	Page   int    `json:"page,omitempty"`
}

type Filter struct {
	Comperator string `json:"comperator,omitempty"`
	Value      string `json:"value,omitempty"`
}
