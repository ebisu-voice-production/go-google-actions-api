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

func (req *AppRequest) GetArgument(name string) string {
	if len(req.Inputs) < 1 {
		return ""
	}
	for _, x := range req.Inputs[0].Arguments {
		if x.Name == name {
			return x.RawText
		}
	}
	return ""
}

func (req *AppRequest) GetConversationToken() string {
	if req.Conversation == nil {
		return ""
	}
	return req.Conversation.ConversationToken
}

func (req *AppRequest) HasSurfaceScreenOutput() bool {
	if req.Surface == nil {
		return false
	}
	for _, capability := range req.Surface.Capabilities {
		if capability.Name == CapabilityScreenOutput {
			return true
		}
	}
	return false
}
