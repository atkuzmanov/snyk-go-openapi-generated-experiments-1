package models

type GetOrgsProjects200Response struct {

	Data []ProjectDeprecated `json:"data"`

	Jsonapi JsonApi `json:"jsonapi"`

	Links Links `json:"links"`
}
