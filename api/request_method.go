package api

import (
	"time"
)

func (req *AppRequest) GetIntent() string {
	if len(req.Inputs) < 1 {
		return ""
	}
	return req.Inputs[0].Intent
}

func (req *AppRequest) GetQuery() string {
	if len(req.Inputs) < 1 {
		return ""
	}
	if len(req.Inputs[0].RawInputs) < 1 {
		return ""
	}
	return req.Inputs[0].RawInputs[0].Query
}

func (req *AppRequest) GetUserId() string {
	if req.User == nil {
		return ""
	}
	return req.User.UserId
}

func (req *AppRequest) GetAccessToken() string {
	if req.User == nil {
		return ""
	}
	return req.User.AccessToken
}

func (req *AppRequest) GetLastSeen() time.Time {
	if req.User == nil {
		return time.Time{}
	}
	return req.User.LastSeen
}

func (req *AppRequest) GetUserStorage() string {
	if req.User == nil {
		return ""
	}
	return req.User.UserStorage
}

func (req *AppRequest) pickFirstArgument(name string) *Argument {
	if len(req.Inputs) < 1 {
		return nil
	}
	for _, x := range req.Inputs[0].Arguments {
		if x.Name == name {
			return &x
		}
	}
	return nil
}

func (req *AppRequest) GetArgument(name string) string {
	argument := req.pickFirstArgument(name)
	if argument == nil {
		return ""
	}
	return argument.RawText
}

func (req *AppRequest) GetArgumentTextValue(name string) string {
	argument := req.pickFirstArgument(name)
	if argument == nil {
		return ""
	}
	return argument.TextValue
}

func (req *AppRequest) GetArgumentBoolValue(name string) bool {
	argument := req.pickFirstArgument(name)
	if argument == nil {
		return false
	}
	return argument.BoolValue
}

func (req *AppRequest) GetMediaStatusArgument() string {
	argument := req.pickFirstArgument("MEDIA_STATUS")
	if argument == nil {
		return ""
	}
	if argument.Extension == nil {
		return ""
	}
	dic, _ := argument.Extension.(map[string]interface{})
	if dic["@type"] != "type.googleapis.com/google.actions.v2.MediaStatus" {
		return ""
	}
	status, _ := dic["status"].(string)
	return status
}

func (req *AppRequest) GetNewSurfaceArgument() string {
	argument := req.pickFirstArgument("NEW_SURFACE")
	if argument == nil {
		return ""
	}
	if argument.Extension == nil {
		return ""
	}
	dic, _ := argument.Extension.(map[string]interface{})
	if dic["@type"] != "type.googleapis.com/google.actions.v2.NewSurfaceValue" {
		return ""
	}
	status, _ := dic["status"].(string)
	return status
}

func (req *AppRequest) GetConversationToken() string {
	if req.Conversation == nil {
		return ""
	}
	return req.Conversation.ConversationToken
}

func (req *AppRequest) hasSurfaceCapability(surfaceCapability SurfaceCapability) bool {
	if req.Surface == nil {
		return false
	}
	for _, capability := range req.Surface.Capabilities {
		if capability.Name == surfaceCapability {
			return true
		}
	}
	return false
}

func (req *AppRequest) hasAvailableSurfaceCapability(surfaceCapability SurfaceCapability) bool {
	for _, surface := range req.AvailableSurfaces {
		for _, capability := range surface.Capabilities {
			if capability.Name == surfaceCapability {
				return true
			}
		}
	}
	return false
}

func (req *AppRequest) HasSurfaceScreenOutput() bool {
	return req.hasSurfaceCapability(CapabilityScreenOutput)
}

func (req *AppRequest) HasAvailableSurfaceScreenOutput() bool {
	return req.hasAvailableSurfaceCapability(CapabilityScreenOutput)
}

func (req *AppRequest) HasSurfaceWebBrowser() bool {
	return req.hasSurfaceCapability(CapabilityWebBrowser)
}

func (req *AppRequest) HasAvailableSurfaceWebBrowser() bool {
	return req.hasAvailableSurfaceCapability(CapabilityWebBrowser)
}

func (req *AppRequest) HasSurfaceMediaResponseAudio() bool {
	return req.hasSurfaceCapability(CapabilityMediaResponseAudio)
}

func (req *AppRequest) HasAvailableSurfaceMediaResponseAudio() bool {
	return req.hasAvailableSurfaceCapability(CapabilityMediaResponseAudio)
}
