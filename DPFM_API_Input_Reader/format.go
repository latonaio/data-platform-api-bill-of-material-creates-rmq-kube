package dpfm_api_input_reader

import (
	"data-platform-api-bill-of-material-creates-rmq-kube/DPFM_API_Caller/requests"
)

func (sdc *SDC) ConvertToHeader() *requests.Header {
	data := sdc.BillOfMaterial
	return &requests.Header{
		PartnerFunction:               data.PartnerFunction,
		BusinessPartner:               data.BusinessPartner,
		BillOfMaterial:                data.BillOfMaterial,
		BillOfMaterialCategory:        data.BillOfMaterialCategory,
		BillOfMaterialVariant:         data.BillOfMaterialVariant,
		BillOfMaterialVersion:         data.BillOfMaterialVersion,
		EngineeringChangeDocument:     data.EngineeringChangeDocument,
		Product:                       data.Product,
		Plant:                         data.Plant,
		BillOfMaterialHeaderUUID:      data.BillOfMaterialHeaderUUID,
		BillOfMaterialVariantUsage:    data.BillOfMaterialVariantUsage,
		EngineeringChangeDocForEdit:   data.EngineeringChangeDocForEdit,
		IsMultipleBOMAlt:              data.IsMultipleBOMAlt,
		BOMHeaderInternalChangeCount:  data.BOMHeaderInternalChangeCount,
		BOMUsagePriority:              data.BOMUsagePriority,
		BillOfMaterialAuthsnGrp:       data.BillOfMaterialAuthsnGrp,
		BOMVersionStatus:              data.BOMVersionStatus,
		IsVersionBillOfMaterial:       data.IsVersionBillOfMaterial,
		IsLatestBOMVersion:            data.IsLatestBOMVersion,
		IsConfiguredMaterial:          data.IsConfiguredMaterial,
		BOMTechnicalType:              data.BOMTechnicalType,
		BOMGroup:                      data.BOMGroup,
		BOMHeaderText:                 data.BOMHeaderText,
		BOMAlternativeText:            data.BOMAlternativeText,
		BillOfMaterialStatus:          data.BillOfMaterialStatus,
		HeaderValidityStartDate:       data.HeaderValidityStartDate,
		HeaderValidityEndDate:         data.HeaderValidityEndDate,
		ChgToEngineeringChgDocument:   data.ChgToEngineeringChgDocument,
		IsMarkedForDeletion:           data.IsMarkedForDeletion,
		IsALE:                         data.IsALE,
		MatFromLotSizeQuantity:        data.MatFromLotSizeQuantity,
		MaterialToLotSizeQuantity:     data.MaterialToLotSizeQuantity,
		BOMHeaderBaseUnit:             data.BOMHeaderBaseUnit,
		BOMHeaderQuantityInBaseUnit:   data.BOMHeaderQuantityInBaseUnit,
		RecordCreationDate:            data.RecordCreationDate,
		LastChangeDate:                data.LastChangeDate,
		BOMIsToBeDeleted:              data.BOMIsToBeDeleted,
		DocumentIsCreatedByCAD:        data.DocumentIsCreatedByCAD,
		LaboratoryOrDesignOffice:      data.LaboratoryOrDesignOffice,
		LastChangeDateTime:            data.LastChangeDateTime,
		ProductDescription:            data.ProductDescription,
		PlantName:                     data.PlantName,
		BillOfMaterialHdrDetailsText:  data.BillOfMaterialHdrDetailsText,
		SelectedBillOfMaterialVersion: data.SelectedBillOfMaterialVersion,
	}
}

func (sdc *SDC) ConvertToItem() *requests.Item {
	dataBillOfMaterial := sdc.BillOfMaterial
	data := sdc.BillOfMaterial.Item
	return &requests.Item{
		BusinessPartner:              dataBillOfMaterial.BusinessPartner,
		BillOfMaterialItemNodeNumber: data.BillOfMaterialItemNodeNumber,
		HeaderChangeDocument:         data.HeaderChangeDocument,
		Material:                     data.Material,
		Plant:                        data.Plant,
		ValidityStartDate:            data.ValidityStartDate,
		ValidityEndDate:              data.ValidityEndDate,
		BillOfMaterialComponent:      data.BillOfMaterialComponent,
		ComponentDescription:         data.ComponentDescription,
		BillOfMaterialItemQuantity:   data.BillOfMaterialItemQuantity,
		ComponentScrapInPercent:      data.ComponentScrapInPercent,
		IsDeleted:                    data.IsDeleted,
	}
}
