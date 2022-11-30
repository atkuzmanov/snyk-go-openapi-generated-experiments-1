package models

import (
	"time"
)

type ProjectDeprecatedAttributes struct {

	BusinessCriticality []string `json:"businessCriticality,omitempty"`

	// The date that the project was created on
	Created time.Time `json:"created"`

	Environment []string `json:"environment,omitempty"`

	Lifecycle []string `json:"lifecycle,omitempty"`

	Name string `json:"name"`

	// The origin the project was added from
	Origin string `json:"origin"`

	// Describes if a project is currently monitored or it is de-activated
	Status string `json:"status"`

	Tags []Tag `json:"tags,omitempty"`

	// The identifier for which revision of the resource is scanned by Snyk. For example this may be a branch for SCM project, or a tag for a container image
	TargetReference *string `json:"targetReference,omitempty"`

	// The package manager of the project
	Type string `json:"type"`
}
