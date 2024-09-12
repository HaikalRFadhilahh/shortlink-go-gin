package requestHandler

type UpdateLinkHanlder struct {
	IsActive *bool   `json:"isActive" binding:"omitempty,boolean"`
	Alias    *string `json:"alias" binding:"omitempty"`
	Link     *string `json:"link" binding:"omitempty"`
}
