package models

type BadRequestError struct {
	Message string `json:"error" example:"invalid input"`
}

func (e BadRequestError) Error() string {
	return e.Message
}

type NotFoundError struct {
	Message string `json:"error" example:"build not found"`
}

func (e NotFoundError) Error() string {
	return e.Message
}

type InternalError struct {
	Message string `json:"error" example:"internal server error"`
}

func (e InternalError) Error() string {
	return e.Message
}
