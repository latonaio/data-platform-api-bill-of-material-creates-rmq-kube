package requests

type Item struct {
	BusinessPartner              *int    `json:"BusinessPartner"`
	BillOfMaterialItemNodeNumber string  `json:"BillOfMaterialItemNodeNumber"`
	HeaderChangeDocument         string  `json:"HeaderChangeDocument"`
	Material                     string  `json:"Material"`
	Plant                        string  `json:"Plant"`
	ValidityStartDate            *string `json:"ValidityStartDate"`
	ValidityEndDate              *string `json:"ValidityEndDate"`
	BillOfMaterialComponent      string  `json:"BillOfMaterialComponent"`
	ComponentDescription         string  `json:"ComponentDescription"`
	BillOfMaterialItemQuantity   string  `json:"BillOfMaterialItemQuantity"`
	ComponentScrapInPercent      string  `json:"ComponentScrapInPercent"`
	IsDeleted                    *bool   `json:"IsDeleted"`
}