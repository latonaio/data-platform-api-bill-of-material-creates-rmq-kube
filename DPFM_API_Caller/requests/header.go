package requests

type Header struct {
	BillOfMaterial                           int     `json:"BillOfMaterial"`
	BillOfMaterialType                       string  `json:"BillOfMaterialType"`
	SupplyChainRelationshipID                int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID        int     `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID   int     `json:"SupplyChainRelationshipDeliveryPlantID"`
	SupplyChainRelationshipProductionPlantID int     `json:"SupplyChainRelationshipProductionPlantID"`
	Product                                  string  `json:"Product"`
	Buyer                                    int     `json:"Buyer"`
	Seller                                   int     `json:"Seller"`
	DestinationDeliverToParty                int     `json:"DestinationDeliverToParty"`
	DestinationDeliverToPlant                string  `json:"DestinationDeliverToPlant"`
	DepartureDeliverFromParty                int     `json:"DepartureDeliverFromParty"`
	DepartureDeliverFromPlant                string  `json:"DepartureDeliverFromPlant"`
	OwnerProductionPlantBusinessPartner      int     `json:"OwnerProductionPlantBusinessPartner"`
	OwnerProductionPlant                     string  `json:"OwnerProductionPlant"`
	ProductBaseUnit                          string  `json:"ProductBaseUnit"`
	ProductDeliveryUnit                      string  `json:"ProductDeliveryUnit"`
	ProductProductionUnit                    string  `json:"ProductProductionUnit"`
	ProductStandardQuantityInBaseUnit        float32 `json:"ProductStandardQuantityInBaseUnit"`
	ProductStandardQuantityInDeliveryUnit    float32 `json:"ProductStandardQuantityInDeliveryUnit"`
	ProductStandardQuantityInProductionUnit  float32 `json:"ProductStandardQuantityInProductionUnit"`
	BillOfMaterialHeaderText                 *string `json:"BillOfMaterialHeaderText"`
	ValidityStartDate                        *string `json:"ValidityStartDate"`
	ValidityEndDate                          *string `json:"ValidityEndDate"`
	CreationDate                             string  `json:"CreationDate"`
	LastChangeDate                           string  `json:"LastChangeDate"`
	IsMarkedForDeletion                      *bool   `json:"IsMarkedForDeletion"`
}
