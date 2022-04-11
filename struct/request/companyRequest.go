package request

type CompanyID struct {
	CompanyID string `json:"companyid" binding:"required"`
}

type CreateCompany struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
