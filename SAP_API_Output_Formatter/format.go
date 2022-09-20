package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-xxxxxxxx-xxxxxxxx-creates-rmq-kube/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToHeader(raw []byte, l *logger.Logger) (*Header, error) {
	pm := &responses.Header{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Header. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := pm.D

	header := &Header{
		XXXXXXXX: data.XXXXXXXX,
	}

	return header, nil
}

func ConvertToItem(raw []byte, l *logger.Logger) (*Item, error) {
	p := &responses.Item{}
	err := json.Unmarshal(raw, p)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to Item. raw data is:\n%v\nunmarshal error: %w", string(raw), err)
	}
	data := p.D
	item := &Item{
		XXXXXXXX: data.XXXXXXXX,
	}

	return item, nil
}
