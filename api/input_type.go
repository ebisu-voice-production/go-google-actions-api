package api

import (
	"bytes"
	"encoding/json"
)

// https://developers.google.com/actions/reference/rest/Shared.Types/InputType

type InputType int

const (
	UnspecifiedInputType InputType = iota
	Touch
	Voice
	Keyboard
)

var inputTypesId = map[InputType]string{
	UnspecifiedInputType: "UNSPECIFIED_INPUT_TYPE",
	Touch:                "TOUCH",
	Voice:                "VOICE",
	Keyboard:             "KEYBOARD",
}

var inputTypesName = map[string]InputType{
	"UNSPECIFIED_INPUT_TYPE": UnspecifiedInputType,
	"TOUCH":                  Touch,
	"VOICE":                  Voice,
	"KEYBOARD":               Keyboard,
}

func (t *InputType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(inputTypesId[*t])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (t *InputType) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*t = inputTypesName[s]
	return nil
}
