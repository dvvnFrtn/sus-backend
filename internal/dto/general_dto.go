package dto

type ResponseID struct {
	ID string `json:"id"`
}

func NewResponseID(id string) *ResponseID {
	return &ResponseID{
		ID: id,
	}
}

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func NewSuccessResponse(msg string, data interface{}) *Response {
	return &Response{
		Status:  true,
		Message: msg,
		Data:    data,
	}
}

func NewErrorResponse(msg string, data interface{}) *Response {
	return &Response{
		Status:  false,
		Message: msg,
		Data:    data,
	}
}
