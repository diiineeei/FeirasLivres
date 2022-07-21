package utils

type JsonHttpResponseMessage struct {
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Response   interface{} `json:"response"`
}
