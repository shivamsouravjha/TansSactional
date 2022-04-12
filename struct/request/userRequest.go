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
