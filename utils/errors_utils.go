package utils

// Util for json errors returns
type ApplicationError struct {
	Message string `json:"message"`
	StatusCode int `json:"status"`
	Code string `json:"code"`
}