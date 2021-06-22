package models

type Photos struct {
	ID string `json:"id"`
}

type Responsible struct {
	ID           string `json:"id"`
	NameOfStep   string `json:"name_of_step"`
	Organization string `json:"organization"`
	Role         string `json:"role"`
	Comment      string `json:"comment"`
	CreatedAt    int32  `json:"created_at"`
	UpdatedAt    int32  `json:"updated_at"`
}

type CreateUpdateResponsibleRequest struct {
	NameOfStep   string `json:"name_of_step"`
	Organization string `json:"organization"`
	Role         string `json:"role"`
	Comment      string `json:"comment"`
	CreatedAt    int32  `json:"created_at"`
	UpdatedAt    int32  `json:"updated_at"`
}

type UpdateResponsibleRequest struct {
	NameOfStep   string `json:"name_of_step"`
	Organization string `json:"organization"`
	Role         string `json:"role"`
	Comment      string `json:"comment"`
	UpdatedAt    int32  `json:"updated_at"`
}

type GetResponsibleResponse struct {
	Responsible *Responsible `json:"responsible"`
}
type GetAllResponsiblesRequest struct {
	Name  string `json:"name"`
	Page  int32  `json:"page"`
	Limit int32  `json:"limit"`
}
type GetAllResponsiblesResponse struct {
	Responsibles []*Responsible `json:"responsibles"`
	Count        int64          `json:"count"`
}
