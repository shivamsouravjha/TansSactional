package request

type UserID struct {
	UserID string `json:"iduser" binding:"required"`
}
