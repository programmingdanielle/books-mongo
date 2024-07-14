package models

type Response struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

type IDResponse struct {
	ID interface{} `json:"id,omitempty"`
}
