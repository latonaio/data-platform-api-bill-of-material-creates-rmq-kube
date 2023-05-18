package requests

type Item struct {
	BusinessPartner                int      `json:"BusinessPartner"`
	BillOfMaterialItem             int      `json:"BillOfMaterialItem"`
	ValidityStartDate              *string  `json:"ValidityStartDate"`
	ValidityEndDate                *string  `json:"ValidityEndDate"`
	BillOfMaterialComponentProduct *string  `json:"BillOfMaterialComponentProduct"`
	BillOfMaterialItemQuantity     *float32 `json:"BillOfMaterialItemQuantity"`
	ComponentScrapInPercent        *int     `json:"ComponentScrapInPercent"`
	IsMarkedForDeletion            *bool    `json:"IsMarkedForDeletion"`
}
