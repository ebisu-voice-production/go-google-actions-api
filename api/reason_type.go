package api

import (
	"bytes"
	"encoding/json"
)

// https://developers.google.com/actions/reference/rest/Shared.Types/ReasonType

type ReasonType int

const (
	UnknownReason ReasonType = iota
	PaymentDeclined
	Ineligible
	PromoNotApplicable
	UnavailableSlot
)

var reasonTypesId = map[ReasonType]string{
	UnknownReason:      "UNKNOWN",
	PaymentDeclined:    "PAYMENT_DECLINED",
	Ineligible:         "INELIGIBLE",
	PromoNotApplicable: "PROMO_NOT_APPLICABLE",
	UnavailableSlot:    "UNAVAILABLE_SLOT",
}

var reasonTypesName = map[string]ReasonType{
	"UNKNOWN":              UnknownReason,
	"PAYMENT_DECLINED":     PaymentDeclined,
	"INELIGIBLE":           Ineligible,
	"PROMO_NOT_APPLICABLE": PromoNotApplicable,
	"UNAVAILABLE_SLOT":     UnavailableSlot,
}

func (t *ReasonType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(reasonTypesId[*t])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (t *ReasonType) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*t = reasonTypesName[s]
	return nil
}
