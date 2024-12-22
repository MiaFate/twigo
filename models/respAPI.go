package models

type ApiResponse[T any] struct {
	Status  int
	Message string
	Data    T
}
