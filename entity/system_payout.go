package entity

type SystemPayout struct {
	CommonResult
	Data []struct {
		PayOutID   string `json:"pay_out_id"`
		PayOutName string `json:"pay_out_name"`
	}
}
