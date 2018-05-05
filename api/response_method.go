package api

func (res *AppResponse) Tell(text string) *AppResponse {
	res.FinalResponse = &FinalResponse{
		RichResponse: &RichResponse{
			Items: []Item{
				{
					SimpleResponse: &SimpleResponse{
						TextToSpeech: text,
					},
				},
			},
		},
	}
	return res
}

func (res *AppResponse) Ask(text string) *AppResponse {
	res.ExpectUserResponse = true
	res.ExpectedInputs = []ExpectedInput{
		{
			InputPrompt: &InputPrompt{
				RichInitialPrompt: &RichResponse{
					Items: []Item{
						{
							SimpleResponse: &SimpleResponse{
								TextToSpeech: text,
							},
						},
					},
				},
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
		RichResponse: &RichResponse{
			Items: []Item{
				{
					SimpleResponse: &SimpleResponse{
						Ssml: ssml,
					},
				},
			},
		},
	}
	return res
}

func (res *AppResponse) AskSsml(ssml string) *AppResponse {
	res.ExpectUserResponse = true
	res.ExpectedInputs = []ExpectedInput{
		{
			InputPrompt: &InputPrompt{
				RichInitialPrompt: &RichResponse{
					Items: []Item{
						{
							SimpleResponse: &SimpleResponse{
								Ssml: ssml,
							},
						},
					},
				},
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

func (res *AppResponse) AttachCardButton(title string, formatted string, label string, url string) *AppResponse {
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
	if res.FinalResponse != nil &&
		res.FinalResponse.RichResponse != nil &&
		len(res.FinalResponse.RichResponse.Items) >= 1 {
		res.FinalResponse.RichResponse.Items = append(res.FinalResponse.RichResponse.Items, item)
	}
	if len(res.ExpectedInputs) > 0 &&
		res.ExpectedInputs[0].InputPrompt != nil &&
		res.ExpectedInputs[0].InputPrompt.RichInitialPrompt != nil &&
		len(res.ExpectedInputs[0].InputPrompt.RichInitialPrompt.Items) >= 1 {
		res.ExpectedInputs[0].InputPrompt.RichInitialPrompt.Items = append(res.ExpectedInputs[0].InputPrompt.RichInitialPrompt.Items, item)
	}
	return res
}

func (res *AppResponse) AttachLinkOut(title string, url string) *AppResponse {
	linkOut := LinkOutSuggestion{
		DestinationName: title,
		OpenUrlAction: &OpenUrlAction{
			Url: url,
		},
	}
	if res.FinalResponse != nil &&
		res.FinalResponse.RichResponse != nil {
		res.FinalResponse.RichResponse.LinkOutSuggestion = &linkOut
	}
	if len(res.ExpectedInputs) > 0 &&
		res.ExpectedInputs[0].InputPrompt != nil &&
		res.ExpectedInputs[0].InputPrompt.RichInitialPrompt != nil {
		res.ExpectedInputs[0].InputPrompt.RichInitialPrompt.LinkOutSuggestion = &linkOut
	}
	return res
}
