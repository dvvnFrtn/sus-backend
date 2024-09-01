package dto

type ResponseID struct {
	ID string `json:"id"`
}

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type RequestIDs struct {
	IDs []string `json:"ids"`
}
