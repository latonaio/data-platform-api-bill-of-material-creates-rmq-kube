package requests

type Header struct {
	BusinessPartner                          int      `json:"BusinessPartner"`
	BillOfMaterial                           int      `json:"BillOfMaterial"`
	Product                                  *string  `json:"Product"`
	Plant                                    *string  `json:"Plant"`
	IsConfiguredComponentProduct             *bool    `json:"IsConfiguredComponentProduct"`
	BOMAlternativeText                       *string  `json:"BOMAlternativeText"`
	HeaderValidityStartDate                  *string  `json:"HeaderValidityStartDate"`
	HeaderValidityEndDate                    *string  `json:"HeaderValidityEndDate"`
	MinumumDeliveryQuantityInBaseUnit        *float32 `json:"MinumumDeliveryQuantityInBaseUnit"`
	MaximumDeliveryLotSizeQuantityInBaseUnit *float32 `json:"MaximumDeliveryLotSizeQuantityInBaseUnit"`
	BOMHeaderBaseUnit                        *string  `json:"BOMHeaderBaseUnit"`
	BOMHeaderQuantityInBaseUnit              *float32 `json:"BOMHeaderQuantityInBaseUnit"`
	CreationDate                             *string  `json:"CreationDate"`
	LastChangeDate                           *string  `json:"LastChangeDate"`
	HeaderText                               *string  `json:"HeaderText"`
	IsMarkedForDeletion                      *bool    `json:"IsMarkedForDeletion"`
}
