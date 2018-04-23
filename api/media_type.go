package api

import (
	"bytes"
	"encoding/json"
)

// https://developers.google.com/actions/reference/rest/Shared.Types/MediaType

type MediaType int

const (
	MediaTypeUnspecified MediaType = iota
	Audio
)

var mediaTypesId = map[MediaType]string{
	MediaTypeUnspecified: "MEDIA_TYPE_UNSPECIFIED",
	Audio:                "AUDIO",
}

var mediaTypesName = map[string]MediaType{
	"MEDIA_TYPE_UNSPECIFIED": MediaTypeUnspecified,
	"AUDIO":                  Audio,
}

func (t *MediaType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(mediaTypesId[*t])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (t *MediaType) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*t = mediaTypesName[s]
	return nil
}
