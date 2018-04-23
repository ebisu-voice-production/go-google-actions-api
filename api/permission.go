package api

import (
	"bytes"
	"encoding/json"
)

// https://developers.google.com/actions/reference/rest/Shared.Types/Permission

type Permission int

const (
	UnspecifiedPermission Permission = iota
	Name
	DevicePreciseLocation
	DeviceCoarseLocation
	Update
)

var permissionsId = map[Permission]string{
	UnspecifiedPermission: "UNSPECIFIED_PERMISSION",
	Name: "NAME",
	DevicePreciseLocation: "DEVICE_PRECISE_LOCATION",
	DeviceCoarseLocation:  "DEVICE_COARSE_LOCATION",
	Update:                "UPDATE",
}

var permissionsName = map[string]Permission{
	"UNSPECIFIED_PERMISSION": UnspecifiedPermission,
	"NAME": Name,
	"DEVICE_PRECISE_LOCATION": DevicePreciseLocation,
	"DEVICE_COARSE_LOCATION":  DeviceCoarseLocation,
	"UPDATE":                  Update,
}

func (t *Permission) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(permissionsId[*t])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

func (t *Permission) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return err
	}
	*t = permissionsName[s]
	return nil
}
