package api

import (
	"bytes"
	"encoding/json"
)

// https://developers.google.com/actions/reference/rest/Shared.Types/ConversationType

type ConversationType int

const (
	TypeUnspecified ConversationType = iota
	New
	Active
)

var conversationTypesId = map[ConversationType]string{
	TypeUnspecified: "TYPE_UNSPECIFIED",
	New:             "NEW",
	Active:          "ACTIVE",
}

var conversationTypesName = map[string]ConversationType{
	"TYPE_UNSPECIFIED": TypeUnspecified,
	"NEW":              New,
	"ACTIVE":           Active,
}

func (t *ConversationType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(conversationTypesId[*t])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (t *ConversationType) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*t = conversationTypesName[s]
	return nil
}
