package models

import (
	"time"
)

type UserAttributesMembership struct {

	// The date the membership was established.
	CreatedAt time.Time `json:"created_at,omitempty"`

	// Whether the membership is a direct, or indirect membership.
	Strategy string `json:"strategy,omitempty"`
}
