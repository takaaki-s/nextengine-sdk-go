package entity

type Authorize struct {
	CommonResult
	CompanyAppHeader string `json:"company_app_header"`
	CompanyNeID      string `json:"company_ne_id"`
	CompanyName      string `json:"company_name"`
	CompanyKana      string `json:"company_kana"`
	UID              string `json:"uid"`
	PicNeID          string `json:"pic_ne_id"`
	PicName          string `json:"pic_name"`
	PicKana          string `json:"pic_kana"`
	PicMailAddress   string `json:"pic_mail_address"`
}
