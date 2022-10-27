package models

type ServerResponse struct {
	Message     string `json:"message"`
	GivenInput  any    `json:"givenInput,omitempty"`
	Code        int    `json:"code"`
	Data        any    `json:"data"`
	CoolorsLink string `json:"coolorsLink"`
}

type ColorMindsResponse struct {
	Result [][]int `json:"result"`
}
