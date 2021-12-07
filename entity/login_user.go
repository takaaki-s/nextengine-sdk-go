package entity

type LoginUser struct {
	CommonResult
	Data []struct {
		PicAddress1            string `json:"pic_address1"`
		PicAddress2            string `json:"pic_address2"`
		PicBirthDate           string `json:"pic_birth_date"`
		PicChargeShopIDList    string `json:"pic_charge_shop_id_list"`
		PicCreationDate        string `json:"pic_creation_date"`
		PicCreatorID           string `json:"pic_creator_id"`
		PicCreatorName         string `json:"pic_creator_name"`
		PicDeletedFlag         string `json:"pic_deleted_flag"`
		PicEnteringCompanyDate string `json:"pic_entering_company_date"`
		PicID                  string `json:"pic_id"`
		PicKana                string `json:"pic_kana"`
		PicLanguageCode        string `json:"pic_language_code"`
		PicLastModifiedByID    string `json:"pic_last_modified_by_id"`
		PicLastModifiedByName  string `json:"pic_last_modified_by_name"`
		PicLastModifiedDate    string `json:"pic_last_modified_date"`
		PicMailAddress         string `json:"pic_mail_address"`
		PicMobile              string `json:"pic_mobile"`
		PicName                string `json:"pic_name"`
		PicNeID                string `json:"pic_ne_id"`
		PicPhone               string `json:"pic_phone"`
		PicPostName            string `json:"pic_post_name"`
		PicRetirementDate      string `json:"pic_retirement_date"`
		PicSkinColor           string `json:"pic_skin_color"`
		PicSocialInsuranceID   string `json:"pic_social_insurance_id"`
		PicSocialInsuranceName string `json:"pic_social_insurance_name"`
		PicTemplateFlag        string `json:"pic_template_flag"`
		PicZipCode             string `json:"pic_zip_code"`
		UID                    string `json:"uid"`
	}
}
