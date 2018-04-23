package api

import (
	"bytes"
	"encoding/json"
)

// https://developers.google.com/actions/reference/rest/Shared.Types/SkuType

type SkuType int

const (
	SkuTypeUnspecified SkuType = iota
	InApp
	Subscription
	App
)

var skuTypesId = map[SkuType]string{
	SkuTypeUnspecified: "TYPE_UNSPECIFIED",
	InApp:              "IN_APP",
	Subscription:       "SUBSCRIPTION",
	App:                "APP",
}

var skuTypesName = map[string]SkuType{
	"TYPE_UNSPECIFIED": SkuTypeUnspecified,
	"IN_APP":           InApp,
	"SUBSCRIPTION":     Subscription,
	"APP":              App,
}

func (t *SkuType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(skuTypesId[*t])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (t *SkuType) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*t = skuTypesName[s]
	return nil
}
