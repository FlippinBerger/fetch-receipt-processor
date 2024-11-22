package api

type ErrorDecodingBody struct {
	Message string `json:"message"`
}

func (e *ErrorDecodingBody) Error() string {
	return "The receipt is invalid" + e.Message
}

type ErrorNotFound struct {
	Message string `json:"message"`
}

func (e *ErrorNotFound) Error() string {
	return "Receipt not found: " + e.Message
}
