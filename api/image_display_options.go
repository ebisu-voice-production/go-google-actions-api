package api

import (
	"bytes"
	"encoding/json"
)

// https://developers.google.com/actions/reference/rest/Shared.Types/ImageDisplayOptions

type ImageDisplayOptions int

const (
	Default ImageDisplayOptions = iota
	White
	Cropped
)

var imageDisplayOptionssId = map[ImageDisplayOptions]string{
	Default: "DEFAULT",
	White:   "WHITE",
	Cropped: "CROPPED",
}

var imageDisplayOptionssName = map[string]ImageDisplayOptions{
	"DEFAULT": Default,
	"WHITE":   White,
	"CROPPED": Cropped,
}

func (t *ImageDisplayOptions) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(imageDisplayOptionssId[*t])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (t *ImageDisplayOptions) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*t = imageDisplayOptionssName[s]
	return nil
}
