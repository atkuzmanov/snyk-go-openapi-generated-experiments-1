package models

type UserAttributes struct {

	// The email of the user.
	Email string `json:"email,omitempty"`

	Membership UserAttributesMembership `json:"membership,omitempty"`

	// The name of the user.
	Name string `json:"name,omitempty"`

	// The username of the user.
	Username string `json:"username,omitempty"`
}
