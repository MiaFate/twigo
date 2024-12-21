package models

type ApiResponse[T any] struct {
	Status     int
	Message    string
	CustomResp T
}
