package sap_api_input_reader

import (
	"sap-api-integrations-xxxxxxxx-xxxxxxxx-creates-rmq-kube/SAP_API_Caller/requests"
)

func (sdc *SDC) ConvertToHeader() *requests.Header {
	data := sdc.Header
	return &requests.Header{
		XXXXXXXX: data.XXXXXXXX,
	}
}

func (sdc *SDC) ConvertToItem() *requests.Item {
	dataXXXXXXXX := sdc.XXXXXXXX
	data := sdc.XXXXXXXX.XXXXXXXXItem
	return &requests.Item{
		XXXXXXXX: dataHeader.XXXXXXXX,
		XXXXXXXX: data.XXXXXXXX,
	}
}
