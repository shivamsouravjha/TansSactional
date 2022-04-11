package response

import structs "transactional/struct"

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
type GetUserReponse struct {
	Message     Response
	UserDetails *structs.UserDetails
}
