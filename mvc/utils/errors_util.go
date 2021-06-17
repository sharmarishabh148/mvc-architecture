package utils

type ApplicationError struct {
	Message    string `josn:"message"`
	StatusCode int    `josn:"status"`
	Code       string `josn:"code"`
}
