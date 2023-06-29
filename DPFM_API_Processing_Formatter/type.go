package dpfm_api_processing_formatter

type HeaderUpdates struct {
	BillOfMaterial                           int     `json:"BillOfMaterial"`
	ProductStandardQuantityInBaseUnit        float32 `json:"ProductStandardQuantityInBaseUnit"`
	ProductStandardQuantityInDeliveryUnit    float32 `json:"ProductStandardQuantityInDeliveryUnit"`
	ProductStandardQuantityInProductionUnit  float32 `json:"ProductStandardQuantityInProductionUnit"`
	BillOfMaterialHeaderText                 *string `json:"BillOfMaterialHeaderText"`
	ValidityStartDate                        *string `json:"ValidityStartDate"`
	ValidityEndDate                          *string `json:"ValidityEndDate"`
}

type ItemUpdates struct {
	BillOfMaterial                                  int      `json:"BillOfMaterial"`
	BillOfMaterialItem                              int      `json:"BillOfMaterialItem"`
	ComponentProductStandardQuantityInBaseUnuit     float32  `json:"ComponentProductStandardQuantityInBaseUnuit"`
	ComponentProductStandardQuantityInDeliveryUnuit float32  `json:"ComponentProductStandardQuantityInDeliveryUnuit"`
	ComponentProductStandardScrapInPercent          *float32 `json:"ComponentProductStandardScrapInPercent"`
	IsMarkedForBackflush                            *bool    `json:"IsMarkedForBackflush"`
	BillOfMaterialItemText                          *string  `json:"BillOfMaterialItemText"`
	ValidityStartDate                               *string  `json:"ValidityStartDate"`
	ValidityEndDate                                 *string  `json:"ValidityEndDate"`
}
