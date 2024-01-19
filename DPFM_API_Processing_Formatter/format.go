package dpfm_api_processing_formatter

import (
	dpfm_api_input_reader "data-platform-api-bill-of-material-creates-rmq-kube/DPFM_API_Input_Reader"
)

func ConvertToHeaderUpdates(header dpfm_api_input_reader.Header) *HeaderUpdates {
	data := header

	return &HeaderUpdates{
		BillOfMaterial:                          data.BillOfMaterial,
		ProductStandardQuantityInBaseUnit:       *data.ProductStandardQuantityInBaseUnit,
		ProductStandardQuantityInDeliveryUnit:   *data.ProductStandardQuantityInDeliveryUnit,
		ProductStandardQuantityInProductionUnit: *data.ProductStandardQuantityInProductionUnit,
		BillOfMaterialHeaderText:                data.BillOfMaterialHeaderText,
		ValidityStartDate:                       data.ValidityStartDate,
		ValidityEndDate:                         data.ValidityEndDate,
	}
}

func ConvertToItemUpdates(header dpfm_api_input_reader.Header, item dpfm_api_input_reader.Item) *ItemUpdates {
	//dataHeader := header
	data := item

	return &ItemUpdates{
		BillOfMaterial:     							data.BillOfMaterial,
		BillOfMaterialItem: 							data.BillOfMaterialItem,
		ComponentProductStandardQuantityInBaseUnit:     *data.ComponentProductStandardQuantityInBaseUnit,
		ComponentProductStandardQuantityInDeliveryUnit: *data.ComponentProductStandardQuantityInDeliveryUnit,
		ComponentProductStandardScrapInPercent:         data.ComponentProductStandardScrapInPercent,
		IsMarkedForBackflush:                           data.IsMarkedForBackflush,
		BillOfMaterialItemText:                         data.BillOfMaterialItemText,
		ValidityStartDate:                              data.ValidityStartDate,
		ValidityEndDate:                                data.ValidityEndDate,
	}
}

func ConvertToItemPricingElementUpdates(header dpfm_api_input_reader.Header, item dpfm_api_input_reader.Item, itemPricingElement dpfm_api_input_reader.ItemPricingElement) *ItemPricingElementUpdates {
	//dataHeader := header
	data := itemPricingElement

	return &ItemPricingElementUpdates{
		BillOfMaterial:     							data.BillOfMaterial,
		BillOfMaterialItem: 							data.BillOfMaterialItem,
		PricingProcedureCounter:						data.PricingProcedureCounter,
		PricingDate:									data.PricingDate,
		ConditionRateValue:								data.ConditionRateValue,
		ConditionRateValueUnit:							data.ConditionRateValueUnit,
		ConditionScaleQuantity:							data.ConditionScaleQuantity,
		TaxCode:										data.TaxCode,
		ConditionAmount:								data.ConditionAmount,
	}
}
