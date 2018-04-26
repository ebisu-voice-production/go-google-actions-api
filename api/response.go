package api

import (
	"time"
)

// https://developers.google.com/actions/reference/rest/Shared.Types/AppResponse

type AppResponse struct {
	ConversationToken  string             `json:"conversationToken,omitempty"`
	UserStorage        string             `json:"userStorage,omitempty"`
	ResetUserStorage   bool               `json:"resetUserStorage,omitempty"`
	ExpectUserResponse bool               `json:"expectUserResponse"`
	ExpectedInputs     []ExpectedInput    `json:"expectedInputs,omitempty"`
	FinalResponse      *FinalResponse     `json:"finalResponse,omitempty"`
	CustomPushMessage  *CustomPushMessage `json:"customPushMessage,omitempty"`
	IsInSandbox        bool               `json:"isInSandbox,omitempty"`
}

type ExpectedInput struct {
	InputPrompt        *InputPrompt     `json:"inputPrompt,omitempty"`
	PossibleIntents    []ExpectedIntent `json:"possibleIntents,omitempty"`
	SpeechBiasingHints []string         `json:"speechBiasingHints,omitempty"`
}

type InputPrompt struct {
	RichInitialPrompt *RichResponse    `json:"richInitialPrompt,omitempty"`
	NoInputPrompts    []SimpleResponse `json:"noInputPrompts,omitempty"`
}

type RichResponse struct {
	Items             []Item             `json:"items,omitempty"`
	Suggestions       []Suggestion       `json:"suggestions,omitempty"`
	LinkOutSuggestion *LinkOutSuggestion `json:"linkOutSuggestion,omitempty"`
}

type Item struct {
	SimpleResponse     *SimpleResponse     `json:"simpleResponse,omitempty"`
	BasicCard          *BasicCard          `json:"basicCard,omitempty"`
	StructuredResponse *StructuredResponse `json:"structuredResponse,omitempty"`
	MediaResponse      *MediaResponse      `json:"mediaResponse,omitempty"`
	CarouselBrowse     *CarouselBrowse     `json:"carouselBrowse,omitempty"`
}

type BasicCard struct {
	Title               string               `json:"title,omitempty"`
	Subtitle            string               `json:"subtitle,omitempty"`
	FormattedText       string               `json:"formattedText,omitempty"`
	Image               *Image               `json:"image,omitempty"`
	Buttons             []Button             `json:"buttons,omitempty"`
	ImageDisplayOptions *ImageDisplayOptions `json:"imageDisplayOptions,omitempty"`
}

type Image struct {
	Url               string `json:"url,omitempty"`
	AccessibilityText string `json:"accessibilityText,omitempty"`
	Height            int    `json:"height,omitempty"`
	Width             int    `json:"width,omitempty"`
}

type Button struct {
	Title         string         `json:"title,omitempty"`
	OpenUrlAction *OpenUrlAction `json:"openUrlAction,omitempty"`
}

type OpenUrlAction struct {
	Url         string       `json:"url,omitempty"`
	AndroidApp  *AndroidApp  `json:"androidApp,omitempty"`
	UrlTypeHint *UrlTypeHint `json:"urlTypeHint,omitempty"`
}

type AndroidApp struct {
	PackageName string          `json:"packageName,omitempty"`
	Versions    []VersionFilter `json:"versions,omitempty"`
}

type VersionFilter struct {
	MinVersion int `json:"minVersion,omitempty"`
	MaxVersion int `json:"maxVersion,omitempty"`
}

type StructuredResponse struct {
	OrderUpdate *OrderUpdate `json:"orderUpdate,omitempty"`
}

type OrderUpdate struct {
	GoogleOrderId          string      `json:"googleOrderId,omitempty"`
	ActionOrderId          string      `json:"actionOrderId,omitempty"`
	OrderState             *OrderState `json:"orderState,omitempty"`
	OrderManagementActions []Action    `json:"orderManagementActions,omitempty"`
	Receipt                *Receipt    `json:"receipt,omitempty"`
	UpdateTime             time.Time   `json:"updateTime,omitempty"`
	TotalPrice             *Price      `json:"totalPrice,omitempty"`
	LineItemUpdates        *struct {
		String *LineItemUpdate `json:"string,omitempty"`
	} `json:"lineItemUpdates,omitempty"`
	UserNotification *UserNotification `json:"userNotification,omitempty"`
	InfoExtension    interface{}       `json:"infoExtension,omitempty"`

	RejectionInfo    *RejectionInfo    `json:"rejectionInfo,omitempty"`
	CancellationInfo *CancellationInfo `json:"cancellationInfo,omitempty"`
	InTransitInfo    *InTransitInfo    `json:"inTransitInfo,omitempty"`
	FullfillmentInfo *FullfillmentInfo `json:"fulfillmentInfo,omitempty"`
	ReturnInfo       *ReturnInfo       `json:"returnInfo,omitempty"`
}

type OrderState struct {
	State string `json:"state,omitempty"`
	Label string `json:"label,omitempty"`
}

type Action struct {
	Type   *ActionType `json:"type,omitempty"`
	Button *Button     `json:"button,omitempty"`
}

type Receipt struct {
	ConfirmedActionOrderId string `json:"confirmedActionOrderId,omitempty"`
	UserVisibleOrderId     string `json:"userVisibleOrderId,omitempty"`
}

type Price struct {
	Type   *PriceType `json:"type,omitempty"`
	Amount *Money     `json:"amount,omitempty"`
}

type Money struct {
	CurrencyCode string `json:"currencyCode,omitempty"`
	Units        string `json:"units,omitempty"`
	Nanos        int    `json:"nanos,omitempty"`
}

type LineItemUpdate struct {
	OrderState *OrderState `json:"orderState,omitempty"`
	Price      *Price      `json:"price,omitempty"`
	Reason     string      `json:"reason,omitempty"`
	Extension  interface{} `json:"extension,omitempty"`
}

type UserNotification struct {
	Title string `json:"title,omitempty"`
	Text  string `json:"text,omitempty"`
}

type RejectionInfo struct {
	Type   *ReasonType `json:"type,omitempty"`
	Reason string      `json:"reason,omitempty"`
}

type CancellationInfo struct {
	Reason string `json:"reason,omitempty"`
}

type InTransitInfo struct {
	UpdatedTime time.Time `json:"updatedTime,omitempty"`
}

type FullfillmentInfo struct {
	DeliveryTime time.Time `json:"deliveryTime,omitempty"`
}

type ReturnInfo struct {
	Reason string `json:"reason,omitempty"`
}

type MediaResponse struct {
	MediaType    *MediaType    `json:"mediaType,omitempty"`
	MediaObjects []MediaObject `json:"mediaObjects,omitempty"`
}

type MediaObject struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	ContentUrl  string `json:"contentUrl,omitempty"`

	LargeImage *Image `json:"largeImage,omitempty"`
	Icon       *Image `json:"icon,omitempty"`
}

type CarouselBrowse struct {
	Items               []Item               `json:"items,omitempty"`
	ImageDisplayOptions *ImageDisplayOptions `json:"imageDisplayOptions,omitempty"`
}

type Suggestion struct {
	Title string `json:"title,omitempty"`
}

type LinkOutSuggestion struct {
	DestinationName string         `json:"destinationName,omitempty"`
	Url             string         `json:"url,omitempty"`
	OpenUrlAction   *OpenUrlAction `json:"openUrlAction,omitempty"`
}

type SimpleResponse struct {
	TextToSpeech string `json:"textToSpeech,omitempty"`
	Ssml         string `json:"ssml,omitempty"`
	DisplayText  string `json:"displayText,omitempty"`
}

type ExpectedIntent struct {
	Intent         string      `json:"intent,omitempty"`
	InputValueData interface{} `json:"inputValueData,omitempty"`
	ParameterName  string      `json:"parameterName,omitempty"`
}

type FinalResponse struct {
	RichResponse *RichResponse `json:"richResponse,omitempty"`
}

type CustomPushMessage struct {
	Target *Target `json:"target,omitempty"`

	OrderUpdate      *OrderUpdate      `json:"orderUpdate,omitempty"`
	UserNotification *UserNotification `json:"userNotification,omitempty"`
}

type Target struct {
	UserId   string    `json:"userId,omitempty"`
	Intent   string    `json:"intent,omitempty"`
	Argument *Argument `json:"argument,omitempty"`
	Locale   string    `json:"locale,omitempty"`
}

func (res *AppResponse) Tell(text string) {
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
}

func (res *AppResponse) Ask(text string) {
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
}
