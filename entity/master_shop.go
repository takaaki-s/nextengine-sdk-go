package entity

type MasterShop struct {
	CommonResult
	Data []MasterShopData
}

type MasterShopData struct {
	ShopAbbreviatedName            string `json:"shop_abbreviated_name"`
	ShopAuthorizationTypeID        string `json:"shop_authorization_type_id"`
	ShopAuthorizationTypeName      string `json:"shop_authorization_type_name"`
	ShopCloseDate                  string `json:"shop_close_date"`
	ShopCreationDate               string `json:"shop_creation_date"`
	ShopCreatorID                  string `json:"shop_creator_id"`
	ShopCreatorName                string `json:"shop_creator_name"`
	ShopCurrencyUnitID             string `json:"shop_currency_unit_id"`
	ShopCurrencyUnitName           string `json:"shop_currency_unit_name"`
	ShopDeletedFlag                string `json:"shop_deleted_flag"`
	ShopHandlingGoodsName          string `json:"shop_handling_goods_name"`
	ShopID                         string `json:"shop_id"`
	ShopKana                       string `json:"shop_kana"`
	ShopLastModifiedByID           string `json:"shop_last_modified_by_id"`
	ShopLastModifiedByName         string `json:"shop_last_modified_by_name"`
	ShopLastModifiedByNullSafeID   string `json:"shop_last_modified_by_null_safe_id"`
	ShopLastModifiedByNullSafeName string `json:"shop_last_modified_by_null_safe_name"`
	ShopLastModifiedDate           string `json:"shop_last_modified_date"`
	ShopLastModifiedNullSafeDate   string `json:"shop_last_modified_null_safe_date"`
	ShopMallID                     string `json:"shop_mall_id"`
	ShopName                       string `json:"shop_name"`
	ShopNote                       string `json:"shop_note"`
	ShopTaxCalculationSequenceID   string `json:"shop_tax_calculation_sequence_id"`
	ShopTaxID                      string `json:"shop_tax_id"`
	ShopTaxName                    string `json:"shop_tax_name"`
	ShopTypeID                     string `json:"shop_type_id"`
}
