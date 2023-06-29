package requests

type Item struct {
	BillOfMaterial                                  int      `json:"BillOfMaterial"`
	BillOfMaterialItem                              int      `json:"BillOfMaterialItem"`
	SupplyChainRelationshipID                       int      `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID               int      `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID          int      `json:"SupplyChainRelationshipDeliveryPlantID"`
	SupplyChainRelationshipStockConfPlantID         int      `json:"SupplyChainRelationshipStockConfPlantID"`
	Product                                         string   `json:"Product"`
	ProductionPlantBusinessPartner                  int      `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                                 string   `json:"ProductionPlant"`
	ComponentProduct                                string   `json:"ComponentProduct"`
	ComponentProductBuyer                           int      `json:"ComponentProductBuyer"`
	ComponentProductSeller                          int      `json:"ComponentProductSeller"`
	ComponentProductDeliverToParty                  int      `json:"ComponentProductDeliverToParty"`
	ComponentProductDeliverToPlant                  string   `json:"ComponentProductDeliverToPlant"`
	ComponentProductDeliverFromParty                int      `json:"ComponentProductDeliverFromParty"`
	ComponentProductDeliverFromPlant                string   `json:"ComponentProductDeliverFromPlant"`
	StockConfirmationBusinessPartner                int      `json:"StockConfirmationBusinessPartner"`
	StockConfirmationPlant                          string   `json:"StockConfirmationPlant"`
	ComponentProductStandardQuantityInBaseUnuit     float32  `json:"ComponentProductStandardQuantityInBaseUnuit"`
	ComponentProductStandardQuantityInDeliveryUnuit float32  `json:"ComponentProductStandardQuantityInDeliveryUnuit"`
	ComponentProductBaseUnit                        string   `json:"ComponentProductBaseUnit"`
	ComponentProductDeliveryUnit                    string   `json:"ComponentProductDeliveryUnit"`
	ComponentProductStandardScrapInPercent          *float32 `json:"ComponentProductStandardScrapInPercent"`
	IsMarkedForBackflush                            *bool    `json:"IsMarkedForBackflush"`
	BillOfMaterialItemText                          *string  `json:"BillOfMaterialItemText"`
	ValidityStartDate                               *string  `json:"ValidityStartDate"`
	ValidityEndDate                                 *string  `json:"ValidityEndDate"`
	CreationDate                                    string   `json:"CreationDate"`
	LastChangeDate                                  string   `json:"LastChangeDate"`
	IsMarkedForDeletion                             *bool    `json:"IsMarkedForDeletion"`
}
