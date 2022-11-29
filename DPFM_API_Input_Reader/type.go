package dpfm_api_input_reader

type EC_MC struct {
}

type SDC struct {
	ConnectionKey       string         `json:"connection_key"`
	Result              bool           `json:"result"`
	RedisKey            string         `json:"redis_key"`
	Filepath            string         `json:"filepath"`
	APIStatusCode       int            `json:"api_status_code"`
	RuntimeSessionID    string         `json:"runtime_session_id"`
	BusinessPartnerID   *int           `json:"business_partner"`
	ServiceLabel        string         `json:"service_label"`
	BillOfMaterial      BillOfMaterial `json:"BillOfMaterial"`
	APISchema           string         `json:"api_schema"`
	Accepter            []string       `json:"accepter"`
	OrderID             *int           `json:"order_id"`
	Deleted             bool           `json:"deleted"`
	SQLUpdateResult     *bool          `json:"sql_update_result"`
	SQLUpdateError      string         `json:"sql_update_error"`
	SubfuncResult       *bool          `json:"subfunc_result"`
	SubfuncError        string         `json:"subfunc_error"`
	ExconfResult        *bool          `json:"exconf_result"`
	ExconfError         string         `json:"exconf_error"`
	APIProcessingResult *bool          `json:"api_processing_result"`
	APIProcessingError  string         `json:"api_processing_error"`
}

type BillOfMaterial struct {
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
	Item                          Item    `json:"Item"`
}

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
