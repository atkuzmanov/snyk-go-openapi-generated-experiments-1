package models

type TargetAttributes struct {

	// A human readable name for this target
	DisplayName string `json:"displayName"`

	// If the resource is private, or publicly accessible
	IsPrivate bool `json:"isPrivate"`

	// The origin that this target relates to
	Origin string `json:"origin"`

	// The url for the resource. Currently only set for targets imported from the CLI
	RemoteUrl *string `json:"remoteUrl"`
}