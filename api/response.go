package api

// https://developers.google.com/actions/reference/rest/Shared.Types/AppResponse

type AppResponse struct {
	ConversationToken  string            `json:"conversationToken,omitempty"`
	UserStorage        string            `json:"userStorage,omitempty"`
	ResetUserStorage   bool              `json:"resetUserStorage,omitempty"`
	ExpectUserResponse bool              `json:"expectUserResponse,omitempty"`
	ExpectedInputs     []ExpectedInput   `json:"expectedInputs,omitempty"`
	FinalResponse      FinalResponse     `json:"finalResponse,omitempty"`
	CustomPushMessage  CustomPushMessage `json:"customPushMessage,omitempty"`
	IsInSandbox        bool              `json:"isInSandbox,omitempty"`
}

type ExpectedInput struct {
	InputPrompt        InputPrompt      `json:"inputPrompt"`
	PossibleIntents    []ExpectedIntent `json:"possibleIntents"`
	SpeechBiasingHints []string         `json:"speechBiasingHints"`
}

type InputPrompt struct {
	RichInitialPrompt RichResponse     `json:"richInitialPrompt"`
	NoInputPrompts    []SimpleResponse `json:"noInputPrompts"`
}

type RichResponse struct {
	Items             []Item            `json:"items"`
	Suggestions       []Suggestion      `json:"suggestions"`
	LinkOutSuggestion LinkOutSuggestion `json:"linkOutSuggestion"`
}

type Item struct {
	SimpleResponse     SimpleResponse     `json:"simpleResponse"`
	BasicCard          BasicCard          `json:"basicCard"`
	StructuredResponse StructuredResponse `json:"structuredResponse"`
	MediaResponse      MediaResponse      `json:"mediaResponse"`
	CarouselBrowse     CarouselBrowse     `json:"carouselBrowse"`
}

type BasicCard struct {
	Title               string              `json:"title"`
	Subtitle            string              `json:"subtitle"`
	FormattedText       string              `json:"formattedText"`
	Image               Image               `json:"image"`
	Buttons             []Button            `json:"buttons"`
	ImageDisplayOptions ImageDisplayOptions `json:"imageDisplayOptions"`
}

type Image struct {
	Url               string `json:"url"`
	AccessibilityText string `json:"accessibilityText"`
	Height            int    `json:"height"`
	Width             int    `json:"width"`
}

type Button struct {
	Title         string        `json:"title"`
	OpenUrlAction OpenUrlAction `json:"openUrlAction"`
}

type OpenUrlAction struct {
	Url         string      `json:"url"`
	AndroidApp  AndroidApp  `json:"androidApp"`
	UrlTypeHint UrlTypeHint `json:"urlTypeHint"`
}

type AndroidApp struct {
	PackageName string          `json:"packageName"`
	Versions    []VersionFilter `json:"versions"`
}

type VersionFilter struct {
	MinVersion int `json:"minVersion"`
	MaxVersion int `json:"maxVersion"`
}

type StructuredResponse struct {
	OrderUpdate OrderUpdate `json:"orderUpdate"`
}

type OrderUpdate struct {
	GoogleOrderId          string     `json:"googleOrderId"`
	ActionOrderId          string     `json:"actionOrderId"`
	OrderState             OrderState `json:"orderState"`
	OrderManagementActions []Action   `json:"orderManagementActions"`
	Receipt                Receipt    `json:"receipt"`
	UpdateTime             string     `json:"updateTime"`
	TotalPrice             Price      `json:"totalPrice"`
	LineItemUpdates        struct {
		String LineItemUpdate `json:"string"`
	} `json:"lineItemUpdates"`
	UserNotification UserNotification `json:"userNotification"`
	InfoExtension    interface{}      `json:"infoExtension"`

	RejectionInfo    RejectionInfo    `json:"rejectionInfo"`
	CancellationInfo CancellationInfo `json:"cancellationInfo"`
	InTransitInfo    InTransitInfo    `json:"inTransitInfo"`
	FullfillmentInfo FullfillmentInfo `json:"fulfillmentInfo"`
	ReturnInfo       ReturnInfo       `json:"returnInfo"`
}

type OrderState struct {
	State string `json:"state"`
	Label string `json:"label"`
}

type Action struct {
	Type   ActionType `json:"type"`
	Button Button     `json:"button"`
}

type Receipt struct {
	ConfirmedActionOrderId string `json:"confirmedActionOrderId"`
	UserVisibleOrderId     string `json:"userVisibleOrderId"`
}

type Price struct {
	Type   PriceType `json:"type"`
	Amount Money     `json:"amount"`
}

type Money struct {
	CurrencyCode string `json:"currencyCode"`
	Units        string `json:"units"`
	Nanos        int    `json:"nanos"`
}

type LineItemUpdate struct {
	OrderState OrderState  `json:"orderState"`
	Price      Price       `json:"price"`
	Reason     string      `json:"reason"`
	Extension  interface{} `json:"extension"`
}

type UserNotification struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type RejectionInfo struct {
	Type   ReasonType `json:"type"`
	Reason string     `json:"reason"`
}

type CancellationInfo struct {
	Reason string `json:"reason"`
}

type InTransitInfo struct {
	UpdatedTime string `json:"updatedTime"`
}

type FullfillmentInfo struct {
	DeliveryTime string `json:"deliveryTime"`
}

type ReturnInfo struct {
	Reason string `json:"reason"`
}

type MediaResponse struct {
	MediaType    MediaType     `json:"mediaType"`
	MediaObjects []MediaObject `json:"mediaObjects"`
}

type MediaObject struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ContentUrl  string `json:"contentUrl"`

	LargeImage Image `json:"largeImage"`
	Icon       Image `json:"icon"`
}

type CarouselBrowse struct {
	Items               []Item              `json:"items"`
	ImageDisplayOptions ImageDisplayOptions `json:"imageDisplayOptions"`
}

type Suggestion struct {
	Title string `json:"title"`
}

type LinkOutSuggestion struct {
	DestinationName string        `json:"destinationName"`
	Url             string        `json:"url"`
	OpenUrlAction   OpenUrlAction `json:"openUrlAction"`
}

type SimpleResponse struct {
	TextToSpeech string `json:"textToSpeech"`
	Ssml         string `json:"ssml"`
	DisplayText  string `json:"displayText"`
}

type ExpectedIntent struct {
	Intent         string      `json:"intent"`
	InputValueData interface{} `json:"inputValueData"`
	ParameterName  string      `json:"parameterName"`
}

type FinalResponse struct {
	RichResponse RichResponse `json:"richResponse"`
}

type CustomPushMessage struct {
	Target Target `json:"target"`

	OrderUpdate      OrderUpdate      `json:"orderUpdate"`
	UserNotification UserNotification `json:"userNotification"`
}

type Target struct {
	UserId   string   `json:"userId"`
	Intent   string   `json:"intent"`
	Argument Argument `json:"argument"`
	Locale   string   `json:"locale"`
}
