package entity

type SystemPaymentMethod struct {
	CommonResult
	Data []struct {
		PaymentMethodID          string `json:"payment_method_id"`
		PaymentMethodName        string `json:"payment_method_name"`
		PaymentMethodDeletedFlag string `json:"payment_method_deleted_flag"`
	}
}
