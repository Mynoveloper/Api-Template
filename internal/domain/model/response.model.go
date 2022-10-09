package model

type Response struct {
	Ok      bool        `json:"ok,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}
