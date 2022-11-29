package requests

type Header struct {
	PartnerFunction               *string `json:"PartnerFunction"`
	BusinessPartner               *int    `json:"BusinessPartner"`
	BillOfMaterial                *string `json:"BillOfMaterial"`
	BillOfMaterialCategory        *string `json:"BillOfMaterialCategory"`
	BillOfMaterialVariant         *string `json:"BillOfMaterialVariant"`
	BillOfMaterialVersion         *string `json:"BillOfMaterialVersion"`
	EngineeringChangeDocument     *string `json:"EngineeringChangeDocument"`
	Product                       *string `json:"Product"`
	Plant                         *string `json:"Plant"`
	BillOfMaterialHeaderUUID      *string `json:"BillOfMaterialHeaderUUID"`
	BillOfMaterialVariantUsage    *string `json:"BillOfMaterialVariantUsage"`
	EngineeringChangeDocForEdit   *string `json:"EngineeringChangeDocForEdit"`
	IsMultipleBOMAlt              *bool   `json:"IsMultipleBOMAlt"`
	BOMHeaderInternalChangeCount  *string `json:"BOMHeaderInternalChangeCount"`
	BOMUsagePriority              *string `json:"BOMUsagePriority"`
	BillOfMaterialAuthsnGrp       *string `json:"BillOfMaterialAuthsnGrp"`
	BOMVersionStatus              *string `json:"BOMVersionStatus"`
	IsVersionBillOfMaterial       *bool   `json:"IsVersionBillOfMaterial"`
	IsLatestBOMVersion            *bool   `json:"IsLatestBOMVersion"`
	IsConfiguredMaterial          *bool   `json:"IsConfiguredMaterial"`
	BOMTechnicalType              *string `json:"BOMTechnicalType"`
	BOMGroup                      *string `json:"BOMGroup"`
	BOMHeaderText                 *string `json:"BOMHeaderText"`
	BOMAlternativeText            *string `json:"BOMAlternativeText"`
	BillOfMaterialStatus          *string `json:"BillOfMaterialStatus"`
	HeaderValidityStartDate       *string `json:"HeaderValidityStartDate"`
	HeaderValidityEndDate         *string `json:"HeaderValidityEndDate"`
	ChgToEngineeringChgDocument   *string `json:"ChgToEngineeringChgDocument"`
	IsMarkedForDeletion           *bool   `json:"IsMarkedForDeletion"`
	IsALE                         *bool   `json:"IsALE"`
	MatFromLotSizeQuantity        *string `json:"MatFromLotSizeQuantity"`
	MaterialToLotSizeQuantity     *string `json:"MaterialToLotSizeQuantity"`
	BOMHeaderBaseUnit             *string `json:"BOMHeaderBaseUnit"`
	BOMHeaderQuantityInBaseUnit   *string `json:"BOMHeaderQuantityInBaseUnit"`
	RecordCreationDate            *string `json:"RecordCreationDate"`
	LastChangeDate                *string `json:"LastChangeDate"`
	BOMIsToBeDeleted              *string `json:"BOMIsToBeDeleted"`
	DocumentIsCreatedByCAD        *bool   `json:"DocumentIsCreatedByCAD"`
	LaboratoryOrDesignOffice      *string `json:"LaboratoryOrDesignOffice"`
	LastChangeDateTime            *string `json:"LastChangeDateTime"`
	ProductDescription            *string `json:"ProductDescription"`
	PlantName                     *string `json:"PlantName"`
	BillOfMaterialHdrDetailsText  *string `json:"BillOfMaterialHdrDetailsText"`
	SelectedBillOfMaterialVersion *string `json:"SelectedBillOfMaterialVersion"`
}
