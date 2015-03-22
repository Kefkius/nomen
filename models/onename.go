package models

// OneName models fields that may be part of the JSON value of a OneName.
type OneName struct {
	Name       interface{}       `json:"name"`
	Avatar     map[string]string `json:"avatar"`
	Cover      map[string]string `json:"cover"`
	Location   interface{}       `json:"location"`
	Website    string            `json:"website"`
	Bio        string            `json:"bio"`
	Bitcoin    interface{}       `json:"bitcoin"`
	Pgp        interface{}       `json:"pgp"`
	Twitter    interface{}       `json:"twitter"`
	Github     interface{}       `json:"github"`
	Facebook   interface{}       `json:"facebook"`
	LinkedIn   interface{}       `json:"linkedin"`
	Bitmessage interface{}       `json:"bitmessage"`
	Version    string            `json:"v"`
}
