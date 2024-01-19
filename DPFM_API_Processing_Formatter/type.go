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
	ComponentProductStandardQuantityInBaseUnit      float32  `json:"ComponentProductStandardQuantityInBaseUnit"`
	ComponentProductStandardQuantityInDeliveryUnit  float32  `json:"ComponentProductStandardQuantityInDeliveryUnit"`
	ComponentProductStandardScrapInPercent          *float32 `json:"ComponentProductStandardScrapInPercent"`
	IsMarkedForBackflush                            *bool    `json:"IsMarkedForBackflush"`
	BillOfMaterialItemText                          *string  `json:"BillOfMaterialItemText"`
	ValidityStartDate                               *string  `json:"ValidityStartDate"`
	ValidityEndDate                                 *string  `json:"ValidityEndDate"`
}

type ItemPricingElementUpdates struct {
	BillOfMaterial				int      `json:"BillOfMaterial"`
	BillOfMaterialItem			int      `json:"BillOfMaterialItem"`
	PricingProcedureCounter		int      `json:"PricingProcedureCounter"`
	PricingDate                 string   `json:"PricingDate"`
	ConditionRateValue          float32  `json:"ConditionRateValue"`
	ConditionRateValueUnit      int      `json:"ConditionRateValueUnit"`
	ConditionScaleQuantity      int      `json:"ConditionScaleQuantity"`
	TaxCode                     *string  `json:"TaxCode"`
	ConditionAmount             float32  `json:"ConditionAmount"`
	CreationDate                string   `json:"CreationDate"`
	LastChangeDate              string   `json:"LastChangeDate"`
	IsCancelled                 *bool    `json:"IsCancelled"`
	IsMarkedForDeletion         *bool    `json:"IsMarkedForDeletion"`
}
