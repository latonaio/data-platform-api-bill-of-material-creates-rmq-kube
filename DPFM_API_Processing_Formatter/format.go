package dpfm_api_processing_formatter

import (
	dpfm_api_input_reader "data-platform-api-bill-of-material-creates-rmq-kube/DPFM_API_Input_Reader"
)

func ConvertToHeaderUpdates(header dpfm_api_input_reader.Header) *HeaderUpdates {
	data := header

	return &HeaderUpdates{
		BillOfMaterial:              data.BillOfMaterial,
		BOMAlternativeText:          data.BOMAlternativeText,
		BOMHeaderQuantityInBaseUnit: data.BOMHeaderQuantityInBaseUnit,
		ValidityStartDate:           data.ValidityStartDate,
		ValidityEndDate:             data.ValidityEndDate,
		BillOfMaterialHeaderText:    data.BillOfMaterialHeaderText,
		IsMarkedForDeletion:         data.IsMarkedForDeletion,
	}
}

func ConvertToItemUpdates(header dpfm_api_input_reader.Header, item dpfm_api_input_reader.Item) *ItemUpdates {
	dataHeader := header
	data := item

	return &ItemUpdates{
		BillOfMaterial:            dataHeader.BillOfMaterial,
		BillOfMaterialItem:        data.BillOfMaterialItem,
		BOMAlternativeText:        data.BOMAlternativeText,
		BOMItemQuantityInBaseUnit: data.BOMItemQuantityInBaseUnit,
		ValidityStartDate:         data.ValidityStartDate,
		ValidityEndDate:           data.ValidityEndDate,
		BillOfMaterialItemText:    data.BillOfMaterialItemText,
		IsMarkedForDeletion:       data.IsMarkedForDeletion,
	}
}
