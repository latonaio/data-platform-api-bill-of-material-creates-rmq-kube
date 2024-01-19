package dpfm_api_output_formatter

import (
	dpfm_api_input_reader "data-platform-api-bill-of-material-creates-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_processing_formatter "data-platform-api-bill-of-material-creates-rmq-kube/DPFM_API_Processing_Formatter"
	"encoding/json"

	"golang.org/x/xerrors"
)

func ConvertToHeaderCreates(sdc *dpfm_api_input_reader.SDC) (*Header, error) {
	data := sdc.Header

	header, err := TypeConverter[*Header](data)
	if err != nil {
		return nil, err
	}

	return header, nil
}

func ConvertToItemCreates(sdc *dpfm_api_input_reader.SDC) (*[]Item, error) {
	items := make([]Item, 0)

	for _, data := range sdc.Header.Item {
		item, err := TypeConverter[*Item](data)
		if err != nil {
			return nil, err
		}

		items = append(items, *item)
	}

	return &items, nil
}

func ConvertToItemPricingElementCreates(sdc *dpfm_api_input_reader.SDC) (*[]ItemPricingElement, error) {
	itemPricingElements := make([]ItemPricingElement, 0)

	for _, data := range sdc.Header.Item.ItemPricingElement {
		itemPricingElement, err := TypeConverter[*ItemPricingElement](data)
		if err != nil {
			return nil, err
		}

		itemPricingElements = append(itemPricingElements, *itemPricingElement)
	}

	return &itemPricingElements, nil
}

func ConvertToHeaderUpdates(headerData dpfm_api_input_reader.Header) (*Header, error) {
	data := headerData

	header, err := TypeConverter[*Header](data)
	if err != nil {
		return nil, err
	}

	return header, nil
}

func ConvertToItemUpdates(itemUpdates *[]dpfm_api_processing_formatter.ItemUpdates) (*[]Item, error) {
	items := make([]Item, 0)

	for _, data := range *itemUpdates {
		item, err := TypeConverter[*Item](data)
		if err != nil {
			return nil, err
		}

		items = append(items, *item)
	}

	return &items, nil
}

func ConvertToItemPricingElementUpdates(itemPricingElementUpdates *[]dpfm_api_processing_formatter.ItemPricingElementUpdates) (*[]ItemPricingElement, error) {
	itemPricingElements := make([]ItemPricingElement, 0)

	for _, data := range *itemPricingElementUpdates {
		itemPricingElement, err := TypeConverter[*ItemPricingElement](data)
		if err != nil {
			return nil, err
		}

		itemPricingElements = append(itemPricingElements, *itemPricingElement)
	}

	return &itemPricingElements, nil
}

func TypeConverter[T any](data interface{}) (T, error) {
	var dist T
	b, err := json.Marshal(data)
	if err != nil {
		return dist, xerrors.Errorf("Marshal error: %w", err)
	}
	err = json.Unmarshal(b, &dist)
	if err != nil {
		return dist, xerrors.Errorf("Unmarshal error: %w", err)
	}
	return dist, nil
}
