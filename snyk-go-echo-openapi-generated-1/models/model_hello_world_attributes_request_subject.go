package models

type HelloWorldAttributesRequestSubject struct {

	ClientId string `json:"clientId,omitempty"`

	PublicId string `json:"publicId"`

	Type string `json:"type"`
}
