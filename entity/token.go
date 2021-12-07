package entity

// Token is Structure that represents a api token
type Token struct {
	AccessToken         string `json:"access_token"`
	RefreshToken        string `json:"refresh_token"`
	AccessTokenEndDate  string `json:"access_token_end_date"`
	RefreshTokenEndDate string `json:"refresh_token_end_date"`
}

// TokenValue returing struct refference
func (t *Token) TokenValue() *Token {
	return t
}
