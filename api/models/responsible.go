package models

import "time"

type Photos struct {
	ID string `json:"id"`
}

type Responsible struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Price     int64     `json:"price"`
	Photos    []*Photos `json:"photos"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateResponsibleRequest struct {
	ID     string    `json:"id"`
	Name   string    `json:"name"`
	Price  int64     `json:"price"`
	Photos []*Photos `json:"photos"`
}

type GetResponsibleResponse struct {
	Responsible *Responsible `json:"responsible"`
}
type GetAllResponsiblesRequest struct {
	Name  string `json:"name"`
	Page  int64  `json:"page"`
	Limit int64  `json:"limit"`
}
type GetAllResponsiblesResponse struct {
	Responsibles []*Responsible `json:"responsibles"`
	Count        int64          `json:"count"`
}
