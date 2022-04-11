package response

import structs "transactional/struct"

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
type GetUserResponse struct {
	Message     Response
	UserDetails *structs.UserDetails
}
type GetAllUserResponse struct {
	Message     Response
	UserDetails *[]structs.UserDetails
}
type GetCompanyResponse struct {
	Message        Response
	CompanyDetails *structs.CompanyDetails
}
type GetAllCompanyResponse struct {
	Message        Response
	CompanyDetails *[]structs.CompanyDetails
}
type CreatedUserResponse struct {
	Status Response
	Token  string `json:"token`
}
