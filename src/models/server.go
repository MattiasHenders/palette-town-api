package models

type ServerResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Data    any    `json:"data"`
}

type ColorMindsResponse struct {
	Result [][]int `json:"result"`
}
