package api

import (
	"bytes"
	"encoding/json"
)

// https://developers.google.com/actions/reference/rest/Shared.Types/PriceType

type PriceType int

const (
	UnknownPriceType PriceType = iota
	Estimate
	Actual
)

var priceTypesId = map[PriceType]string{
	UnknownPriceType: "UNKNOWN",
	Estimate:         "ESTIMATE",
	Actual:           "ACTUAL",
}

var priceTypesName = map[string]PriceType{
	"UNKNOWN":  UnknownPriceType,
	"ESTIMATE": Estimate,
	"ACTUAL":   Actual,
}

func (t *PriceType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(priceTypesId[*t])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (t *PriceType) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*t = priceTypesName[s]
	return nil
}
