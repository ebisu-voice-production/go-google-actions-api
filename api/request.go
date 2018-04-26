package api

import (
	"time"
)

// https://developers.google.com/actions/reference/rest/Shared.Types/AppRequest

type AppRequest struct {
	User              *User         `json:"user"`
	Device            *Device       `json:"device"`
	Surface           *Surface      `json:"surface"`
	Conversation      *Conversation `json:"conversation"`
	Inputs            []Input       `json:"inputs"`
	IsInSandbox       bool          `json:"isInSandbox"`
	AvailableSurfaces []Surface     `json:"availableSurfaces"`
}

type User struct {
	UserId              string               `json:"userId"`
	Profile             *UserProfile         `json:"profile"`
	AccessToken         string               `json:"accessToken"`
	Permissions         []Permission         `json:"permissions"`
	Locale              string               `json:"locale"`
	LastSeen            time.Time            `json:"lastSeen"`
	UserStorage         string               `json:"userStorage"`
	PackageEntitlements []PackageEntitlement `json:"packageEntitlements"`
}

type UserProfile struct {
	DisplayName string `json:"displayName"`
	GivenName   string `json:"givenName"`
	FamilyName  string `json:"familyName"`
}

type PackageEntitlement struct {
	PackageName  string        `json:"packageName"`
	Entitlements []Entitlement `json:"entitlements"`
}

type Entitlement struct {
	Sku          string      `json:"sku"`
	SkuType      *SkuType    `json:"skuType"`
	InAppDetails interface{} `json:"inAppDetails"`
}

type Device struct {
	Location *Location `json:"location"`
}

type Location struct {
	Coordinates      *LatLng        `json:"coordinates"`
	FormattedAddress string         `json:"formattedAddress"`
	ZipCode          string         `json:"zipCode"`
	City             string         `json:"city"`
	PostalAddress    *PostalAddress `json:"postalAddress"`
	Name             string         `json:"name"`
	PhoneNumber      string         `json:"phoneNumber"`
	Notes            string         `json:"notes"`
}

type LatLng struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type PostalAddress struct {
	Revision           int      `json:"revision"`
	RegionCode         string   `json:"regionCode"`
	LanguageCode       string   `json:"languageCode"`
	PostalCode         string   `json:"postalCode"`
	SortingCode        string   `json:"sortingCode"`
	AdministrativeArea string   `json:"administrativeArea"`
	Locality           string   `json:"locality"`
	Sublocality        string   `json:"sublocality"`
	AddressLines       []string `json:"addressLines"`
	Recipients         []string `json:"recipients"`
	Organization       string   `json:"organization"`
}

type Surface struct {
	Capabilities []Capability `json:"capabilities"`
}

type Capability struct {
	Name string `json:"name"`
}

type Conversation struct {
	ConversationId    string            `json:"conversationId"`
	Type              *ConversationType `json:"type"`
	ConversationToken string            `json:"conversationToken"`
}

type Input struct {
	RawInputs []RawInput `json:"rawInputs"`
	Intent    string     `json:"intent"`
	Arguments []Argument `json:"arguments"`
}

type RawInput struct {
	InputType *InputType `json:"inputType"`
	Query     string     `json:"query"`
}

type Argument struct {
	Name      string  `json:"name"`
	RawText   string  `json:"rawText"`
	TextValue string  `json:"textValue"`
	Status    *Status `json:"status"`

	IntValue        string      `json:"intValue"` // Is this a bug in the spec?
	FloatValue      float64     `json:"floatValue"`
	BoolValue       bool        `json:"boolValue"`
	DatetimeValue   *DateTime   `json:"datetimeValue"`
	PlaceValue      *Location   `json:"placeValue"`
	Extension       interface{} `json:"extension"`
	StructuredValue interface{} `json:"structuredValue"`
}

type Status struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details"`
}

type DateTime struct {
	Date *Date      `json:"date"`
	Time *TimeOfDay `json:"time"`
}

type Date struct {
	Year  int `json:"year"`
	Month int `json:"month"`
	Day   int `json:"day"`
}

type TimeOfDay struct {
	Hours   int `json:"hours"`
	Minutes int `json:"minutes"`
	Seconds int `json:"seconds"`
	Nanos   int `json:"nanos"`
}

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
