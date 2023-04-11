package models

type Response struct {
	Status string `json:"status"`
}

type ErrorResponse struct {
	Response
	Code  int    `json:"code"`
	Error string `json:"error"`
}
