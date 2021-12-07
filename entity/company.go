package entity

type Company struct {
	CommonResult
	Data []struct {
		CompanyID string `json:"company_id"`
	}
}
