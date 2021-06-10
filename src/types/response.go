package types

type Response struct {
	Success bool        `json:"success"`
	Error   string      `json:"error,omitempty"`
	Code    int64       `json:"code,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func NewResponseSuccess(code int64, payload interface{}) *Response {
	return &Response{
		true,
		"",
		code,
		payload,
	}
}

func NewResponseError(code int64, err string) *Response {
	return &Response{
		false,
		err,
		code,
		nil,
	}
}
