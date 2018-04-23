package api

import (
	"bytes"
	"encoding/json"
)

// https://developers.google.com/actions/reference/rest/Shared.Types/UrlTypeHint

type UrlTypeHint int

const (
	UrlTypeHintUnspecified UrlTypeHint = iota
	AmpContent
)

var urlTypeHintsId = map[UrlTypeHint]string{
	UrlTypeHintUnspecified: "URL_TYPE_HINT_UNSPECIFIED",
	AmpContent:             "AMP_CONTENT",
}

var urlTypeHintsName = map[string]UrlTypeHint{
	"URL_TYPE_HINT_UNSPECIFIED": UrlTypeHintUnspecified,
	"AMP_CONTENT":               AmpContent,
}

func (t *UrlTypeHint) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(urlTypeHintsId[*t])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (t *UrlTypeHint) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*t = urlTypeHintsName[s]
	return nil
}
