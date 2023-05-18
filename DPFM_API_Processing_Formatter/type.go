package dpfm_api_processing_formatter

type HeaderUpdates struct {
	BillOfMaterial              int      `json:"BillOfMaterial"`
	BOMAlternativeText          *string  `json:"BOMAlternativeText"`
	BOMHeaderQuantityInBaseUnit *float32 `json:"BOMHeaderQuantityInBaseUnit"`
	ValidityStartDate           *string  `json:"ValidityStartDate"`
	ValidityEndDate             *string  `json:"ValidityEndDate"`
	BillOfMaterialHeaderText    *string  `json:"BillOfMaterialHeaderText"`
	IsMarkedForDeletion         *bool    `json:"IsMarkedForDeletion"`
}

type ItemUpdates struct {
	BillOfMaterial            int      `json:"BillOfMaterial"`
	BillOfMaterialItem        int      `json:"BillOfMaterialItem"`
	BOMAlternativeText        *string  `json:"BOMAlternativeText"`
	BOMItemQuantityInBaseUnit *float32 `json:"BOMItemQuantityInBaseUnit"`
	ValidityStartDate         *string  `json:"ValidityStartDate"`
	ValidityEndDate           *string  `json:"ValidityEndDate"`
	BillOfMaterialItemText    *string  `json:"BillOfMaterialItemText"`
	IsMarkedForDeletion       *bool    `json:"IsMarkedForDeletion"`
}
