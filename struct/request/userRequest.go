package request

type UserID struct {
	UserID string `json:"iduser" binding:"required"`
}

type CreateUser struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	CompanyID int    `json:"companyid"`
	Password  string `json:"password"`
}
