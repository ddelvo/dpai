package dto

type Response struct {
	Message any `json:"message,omitempty"`
	Data    any `json:"data,omitempty"`
}
