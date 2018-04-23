package api

import (
	"bytes"
	"encoding/json"
)

// https://developers.google.com/actions/reference/rest/Shared.Types/ActionType

type ActionType int

const (
	UnknownActionType ActionType = iota
	ViewDetails
	Modify
	Cancel
	Return
	Exchange
	Email
	Call
	Reorder
	Review
	CustomerService
)

var actionTypesId = map[ActionType]string{
	UnknownActionType: "UNKNOWN",
	ViewDetails:       "VIEW_DETAILS",
	Modify: "MODIFY	",
	Cancel:   "CANCEL",
	Return:   "RETURN",
	Exchange: "EXCHANGE",
	Email:    "EMAIL",
	Call:     "CALL",
	Reorder:  "REORDER",
	Review: "REVIEW	",
	CostomerSerivce: "CUSTOMER_SERVICE",
}

var actionTypesName = map[string]ActionType{
	"UNKNOWN":      UnknownActionType,
	"VIEW_DETAILS": ViewDetails,
	"MODIFY	": Modify,
	"CANCEL":   Cancel,
	"RETURN":   Return,
	"EXCHANGE": Exchange,
	"EMAIL":    Email,
	"CALL":     Call,
	"REORDER":  Reorder,
	"REVIEW	": Review,
	"CUSTOMER_SERVICE": CostomerSerivce,
}

func (t *ActionType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(actionTypesId[*t])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (t *ActionType) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*t = actionTypesName[s]
	return nil
}
