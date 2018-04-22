package api

// https://developers.google.com/actions/reference/rest/Shared.Types/AppRequest

type AppRequest struct {
	User              User         `json:"user"`
	Device            Device       `json:"device"`
	Surface           Surface      `json:"surface"`
	Conversation      Conversation `json:"conversation"`
	Inputs            []Input      `json:"inputs"`
	IsInSandbox       bool         `json:"isInSandbox"`
	AvailableSurfaces []Surface    `json:"availableSurfaces"`
}

type User struct {
	// TODO
}

type Device struct {
	// TODO
}

type Surface struct {
	// TODO
}

type Conversation struct {
	// TODO
}

type Input struct {
	// TODO
}
