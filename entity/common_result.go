package entity

type CommonResult struct {
	Result  string `json:"result"`
	Count   string `json:"count"`
	Message string `json:"message"`
	Code    string `json:"code"`
	Token
}
