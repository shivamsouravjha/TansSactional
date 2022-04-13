package request

type CompanyID struct {
	CompanyID string `json:"companyid" binding:"required"`
}

type CreateCompany struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateCompany struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
