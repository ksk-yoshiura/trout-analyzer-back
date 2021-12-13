package controllers

type response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

func newResponse(status int, message string, result interface{}) *response {
	return &response{status, message, result}
}
