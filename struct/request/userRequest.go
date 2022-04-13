package request

type UserID struct {
	UserID string `json:"iduser" binding:"required"`
}

type CreateUser struct {
	Name      string `json:"name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	CompanyID int    `json:"companyid" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

type CreateProduct struct {
	ProductName string `json:"productname" binding:"required"`
	Price       int    `json:"price" binding:"required"`
	CompanyID   string `json:"companyid"`
}
