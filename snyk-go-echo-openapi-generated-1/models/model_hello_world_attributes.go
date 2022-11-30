package models

type HelloWorldAttributes struct {

	BetaField string `json:"betaField"`

	Message string `json:"message"`

	RequestSubject HelloWorldAttributesRequestSubject `json:"requestSubject"`
}
