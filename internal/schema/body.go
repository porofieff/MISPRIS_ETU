package schema

type CreateBodyInput struct {
	CarcassID string `json:"carcass_id" binding:"required"`
	DoorsID   string `json:"doors_id" binding:"required"`
	WingsID   string `json:"wings_id" binding:"required"`
}

type CreateCarcassInput struct {
	Name string `json:"name" binding:"required"`
	Info string `json:"info"`
}

type CreateDoorsInput struct {
	Name string `json:"name" binding:"required"`
	Info string `json:"info"`
}

type CreateWingsInput struct {
	Name string `json:"name" binding:"required"`
	Info string `json:"info"`
}
