package api

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

func (req *AppRequest) GetLastSeen() time.Time {
	if req.User == nil {
		return time.Time{}
	}
	return req.User.LastSeen
}
