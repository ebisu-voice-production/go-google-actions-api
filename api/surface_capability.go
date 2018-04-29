package api

import (
	"bytes"
	"encoding/json"
)

// https://developers.google.com/actions/reference/rest/Shared.Types/AppRequest#Capability

type SurfaceCapability int

const (
	CapabilityAudioOutput SurfaceCapability = iota
	CapabilityScreenOutput
	CapabilityMediaResponseAudio
	CapabilityWebBrowser
)

var surfaceCapabilitysId = map[SurfaceCapability]string{
	CapabilityAudioOutput:        "actions.capability.AUDIO_OUTPUT",
	CapabilityScreenOutput:       "actions.capability.SCREEN_OUTPUT",
	CapabilityMediaResponseAudio: "actions.capability.MEDIA_RESPONSE_AUDIO",
	CapabilityWebBrowser:         "actions.capability.WEB_BROWSER",
}

var surfaceCapabilitysName = map[string]SurfaceCapability{
	"actions.capability.AUDIO_OUTPUT":         CapabilityAudioOutput,
	"actions.capability.SCREEN_OUTPUT":        CapabilityScreenOutput,
	"actions.capability.MEDIA_RESPONSE_AUDIO": CapabilityMediaResponseAudio,
	"actions.capability.WEB_BROWSER":          CapabilityWebBrowser,
}

func (t *SurfaceCapability) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(surfaceCapabilitysId[*t])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (t *SurfaceCapability) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*t = surfaceCapabilitysName[s]
	return nil
}
