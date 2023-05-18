package existence_conf

import (
	dpfm_api_input_reader "data-platform-api-bill-of-material-creates-rmq-kube/DPFM_API_Input_Reader"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	"golang.org/x/xerrors"
)

func (c *ExistenceConf) headerPlantExistenceConf(mapper ExConfMapper, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, exconfErrMsg *string, errs *[]error, mtx *sync.Mutex, wg *sync.WaitGroup, log *logger.Logger) {
	defer wg.Done()
	wg2 := sync.WaitGroup{}
	exReqTimes := 0

	headers := make([]dpfm_api_input_reader.Header, 0, 1)
	headers = append(headers, input.Header)
	for _, header := range headers {
		plant, bpID := getHeaderPlantExistenceConfKey(mapper, &header, exconfErrMsg)
		wg2.Add(1)
		exReqTimes++
		go func() {
			if isZero(plant) || isZero(bpID) {
				wg2.Done()
				return
			}
			res, err := c.plantExistenceConfRequest(plant, bpID, mapper, input, existenceMap, mtx, log)
			if err != nil {
				mtx.Lock()
				*errs = append(*errs, err)
				mtx.Unlock()
			}
			if res != "" {
				*exconfErrMsg = res
			}
			wg2.Done()
		}()
	}
	wg2.Wait()
	if exReqTimes == 0 {
		*existenceMap = append(*existenceMap, false)
	}
}

func (c *ExistenceConf) itemPlantExistenceConf(mapper ExConfMapper, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, exconfErrMsg *string, errs *[]error, mtx *sync.Mutex, wg *sync.WaitGroup, log *logger.Logger) {
	defer wg.Done()
	wg2 := sync.WaitGroup{}
	exReqTimes := 0

	items := input.Header.Item
	for _, item := range items {
		plant, bpID := getItemPlantExistenceConfKey(mapper, &item, exconfErrMsg)
		wg2.Add(1)
		exReqTimes++
		go func() {
			if isZero(plant) || isZero(bpID) {
				wg2.Done()
				return
			}
			res, err := c.plantExistenceConfRequest(plant, bpID, mapper, input, existenceMap, mtx, log)
			if err != nil {
				mtx.Lock()
				*errs = append(*errs, err)
				mtx.Unlock()
			}
			if res != "" {
				*exconfErrMsg = res
			}
			wg2.Done()
		}()
	}
	wg2.Wait()
	if exReqTimes == 0 {
		*existenceMap = append(*existenceMap, false)
	}
}

func (c *ExistenceConf) plantExistenceConfRequest(plant string, bpID int, mapper ExConfMapper, input *dpfm_api_input_reader.SDC, existenceMap *[]bool, mtx *sync.Mutex, log *logger.Logger) (string, error) {
	keys := newResult(map[string]interface{}{
		"BusinessPartner": bpID,
		"Plant":           plant,
	})
	exist := false
	defer func() {
		mtx.Lock()
		*existenceMap = append(*existenceMap, exist)
		mtx.Unlock()
	}()

	req, err := jsonTypeConversion[Returns](input)
	if err != nil {
		return "", xerrors.Errorf("request create error: %w", err)
	}
	req.PlantGeneralReturn.Plant = plant
	req.PlantGeneralReturn.BusinessPartner = bpID

	exist, err = c.exconfRequest(req, mapper, log)
	if err != nil {
		return "", err
	}
	if !exist {
		return keys.fail(), nil
	}
	return "", nil
}

func getHeaderPlantExistenceConfKey(mapper ExConfMapper, header *dpfm_api_input_reader.Header, exconfErrMsg *string) (string, int) {
	var plant string
	var bpID int

	switch mapper.Field {
	case "OwnerPlant":
		if header.OwnerPlant == nil || header.OwnerBusinessPartner == nil {
			plant = ""
			bpID = 0
		} else {
			plant = *header.OwnerPlant
			bpID = *header.OwnerBusinessPartner
		}
	}

	return plant, bpID
}

func getItemPlantExistenceConfKey(mapper ExConfMapper, item *dpfm_api_input_reader.Item, exconfErrMsg *string) (string, int) {
	var plant string
	var bpID int

	if item.StockConfirmationPlant == nil || item.ComponentProductBusinessPartner == nil {
		plant = ""
		bpID = 0
	} else {
		plant = *item.StockConfirmationPlant
		bpID = *item.ComponentProductBusinessPartner
	}

	return plant, bpID
}
