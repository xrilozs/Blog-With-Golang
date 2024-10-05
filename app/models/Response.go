package models

type APIResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type TokenResponse struct {
	Token   string `json:"token"`
	Refresh string `json:"refresh"`
}
