package api

func buildSimpleResponse(text string, ssml string) *SimpleResponse {
	var simpleResponse SimpleResponse
	if ssml != "" {
		simpleResponse.Ssml = ssml
		if text != "" {
			simpleResponse.DisplayText = text
		}
	} else {
		simpleResponse.TextToSpeech = text
	}
	return &simpleResponse
}

func buildRichResponse(text string, ssml string) *RichResponse {
	richResponse := RichResponse{
		Items: []Item{
			{
				SimpleResponse: buildSimpleResponse(text, ssml),
			},
		},
	}
	return &richResponse
}

func (res *AppResponse) Tell(text string) *AppResponse {
	res.FinalResponse = &FinalResponse{
		RichResponse: buildRichResponse(text, ""),
	}
	return res
}

func (res *AppResponse) Ask(text string) *AppResponse {
	res.ExpectUserResponse = true
	res.ExpectedInputs = []ExpectedInput{
		{
			InputPrompt: &InputPrompt{
				RichInitialPrompt: buildRichResponse(text, ""),
			},
			PossibleIntents: []ExpectedIntent{
				{
					Intent: "actions.intent.TEXT",
				},
			},
		},
	}
	return res
}

func (res *AppResponse) TellSsml(ssml string) *AppResponse {
	res.FinalResponse = &FinalResponse{
		RichResponse: buildRichResponse("", ssml),
	}
	return res
}

func (res *AppResponse) AskSsml(ssml string) *AppResponse {
	res.ExpectUserResponse = true
	res.ExpectedInputs = []ExpectedInput{
		{
			InputPrompt: &InputPrompt{
				RichInitialPrompt: buildRichResponse("", ssml),
			},
			PossibleIntents: []ExpectedIntent{
				{
					Intent: "actions.intent.TEXT",
				},
			},
		},
	}
	return res
}

func findRichResponse(res *AppResponse) *RichResponse {
	if res.FinalResponse != nil &&
		res.FinalResponse.RichResponse != nil {
		return res.FinalResponse.RichResponse
	}
	if len(res.ExpectedInputs) > 0 &&
		res.ExpectedInputs[0].InputPrompt != nil &&
		res.ExpectedInputs[0].InputPrompt.RichInitialPrompt != nil {
		return res.ExpectedInputs[0].InputPrompt.RichInitialPrompt
	}
	return nil
}

func (res *AppResponse) AttachCardButton(title string, formatted string, label string, url string) *AppResponse {
	richResponse := findRichResponse(res)
	if richResponse == nil {
		return res
	}
	item := Item{
		BasicCard: &BasicCard{
			Title:         title,
			FormattedText: formatted,
			Buttons: []Button{
				{
					Title: label,
					OpenUrlAction: &OpenUrlAction{
						Url: url,
					},
				},
			},
		},
	}
	if len(richResponse.Items) >= 1 {
		richResponse.Items = append(richResponse.Items, item)
	}
	return res
}

func (res *AppResponse) AttachLinkOut(title string, url string) *AppResponse {
	richResponse := findRichResponse(res)
	if richResponse == nil {
		return res
	}
	linkOut := LinkOutSuggestion{
		DestinationName: title,
		OpenUrlAction: &OpenUrlAction{
			Url: url,
		},
	}
	richResponse.LinkOutSuggestion = &linkOut
	return res
}

func (res *AppResponse) AskForConfirmation(text string) *AppResponse {
	inputValueData := InputValueDataForConfirmation{
		Type: "type.googleapis.com/google.actions.v2.ConfirmationValueSpec",
	}
	inputValueData.DialogSpec.RequestConfirmationText = text
	res.ExpectUserResponse = true
	res.ExpectedInputs = []ExpectedInput{
		{
			InputPrompt: &InputPrompt{
				RichInitialPrompt: buildRichResponse("PLACEHOLDER_FOR_CONFIRMATION", ""),
			},
			PossibleIntents: []ExpectedIntent{
				{
					Intent:         "actions.intent.CONFIRMATION",
					InputValueData: inputValueData,
				},
			},
		},
	}
	return res
}

func (res *AppResponse) AskForSignIn() *AppResponse {
	res.ExpectUserResponse = true
	res.ExpectedInputs = []ExpectedInput{
		{
			InputPrompt: &InputPrompt{
				RichInitialPrompt: buildRichResponse("PLACEHOLDER_FOR_SIGN_IN", ""),
			},
			PossibleIntents: []ExpectedIntent{
				{
					Intent: "actions.intent.SIGN_IN",
				},
			},
		},
	}
	return res
}

func (res *AppResponse) askForNewSurface(context string, notification string, surfaceCapability SurfaceCapability) *AppResponse {
	inputValueData := InputValueDataForNewSurface{
		Type:              "type.googleapis.com/google.actions.v2.NewSurfaceValueSpec",
		Context:           context,
		NotificationTitle: notification,
		Capabilities:      []SurfaceCapability{surfaceCapability},
	}
	res.ExpectUserResponse = true
	res.ExpectedInputs = []ExpectedInput{
		{
			InputPrompt: &InputPrompt{
				RichInitialPrompt: buildRichResponse("PLACEHOLDER_FOR_NEW_SURFACE", ""),
			},
			PossibleIntents: []ExpectedIntent{
				{
					Intent:         "actions.intent.NEW_SURFACE",
					InputValueData: inputValueData,
				},
			},
		},
	}
	return res
}

func (res *AppResponse) AskForScreenOutput(context string, notification string) *AppResponse {
	return res.askForNewSurface(context, notification, CapabilityScreenOutput)
}

func (res *AppResponse) AskForWebBrowser(context string, notification string) *AppResponse {
	return res.askForNewSurface(context, notification, CapabilityWebBrowser)
}
