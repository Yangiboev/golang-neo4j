package models

type CreateResponse struct {
	ID string `json:"id"`
}

type GetRequest struct {
	ID string `json:"id"`
}

type DeleteRequest struct {
	ID string `json:"id"`
}

type SuccessResponse struct {
	Success bool `json:"success" example:"true"`
}

//InternalServerError ...
type InternalServerError struct {
	Code    string
	Message string
}

//BadRequest
type BadRequestError struct {
	Code    string
	Message string
}

//ValidationError ...
type ValidationError struct {
	Code        string
	Message     string
	UserMessage string
}
