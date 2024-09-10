package requestHandler

type CreateLinkHeader struct {
	IsActive bool    `json:"isActive" binding:"omitempty"`
	Alias    *string `json:"alias" binding:"required"`
	Link     *string `json:"link" binding:"required"`
	UserId   *int    `json:"userId,omitempty" binding:"omitempty"`
}
