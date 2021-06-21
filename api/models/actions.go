package models

type Action struct {
	Type      string `json:"type"`
	Status    string `json:"status"`
	Comment   string `json:"comment"`
	Role      string `json:"role"`
	CreatedAt int32  `json:"created_at"`
}
