package models

type Action struct {
	ID        string `json:"id"`
	Type      string `json:"type"`
	Status    string `json:"status"`
	Comment   string `json:"comment"`
	Role      string `json:"role"`
	CreatedAt int64  `json:"created_at"`
}

type CreateActionRequest struct {
	Type      string `json:"type"`
	Status    string `json:"status"`
	Comment   string `json:"comment"`
	Role      string `json:"role"`
	CreatedAt int64  `json:"created_at"`
}

type UpdateActionRequest struct {
	Type      string `json:"type"`
	Status    string `json:"status"`
	Comment   string `json:"comment"`
	Role      string `json:"role"`
	CreatedAt int64  `json:"created_at"`
}

type GetAllActionsRequest struct {
	Page  int32 `json:"page"`
	Limit int32 `json:"limit"`
}
type GetAllActionsResponse struct {
	Actions []*Action `json:"action"`
	Count   int64     `json:"count"`
}
