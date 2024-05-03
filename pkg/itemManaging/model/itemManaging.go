package model

type (
	ItemCreatingReq struct {
		AdminID     string
		Name        string `json:"name" validate:"required,max=64"`
		Description string `json:"description" validate:"required,max=128"`
		Picture     string `json:"picture" validate:"required"`
		Price       uint   `json:"price" validate:"required"`
	}

	ItemEditingReq struct {
		AdminID     string
		ID          uint64 `json:"id" validate:"required,max=64"`
		Name        string `json:"name" validate:"required,max=64"`
		Description string `json:"description" validate:"omitempty,max=128"`
		Picture     string `json:"picture" validate:"omitempty"`
		Price       uint   `json:"price" validate:"omitempty"`
	}
)
