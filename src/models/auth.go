package models

type TokenDetails struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	AccessUuid   string `json:"accessID"`
	RefreshUuid  string `json:"refreshID"`
	AtExpires    int64  `json:"accessExpires"`
	RtExpires    int64  `json:"refreshExpires"`
}

type AccessDetails struct {
	AccessUuid string
	UserEmail  string
}
