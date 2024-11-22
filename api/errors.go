package api

type ErrorDecodingBody struct {
	error string
}

func (e *ErrorDecodingBody) Error() string {
	return "The receipt is invalid" + e.error
}
