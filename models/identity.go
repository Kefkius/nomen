package models

import (
	"encoding/json"
)

// IdentityShow models fields in a name_show result
type IdentityShow struct {
	Name        string                 `json:"name"`
	Value       Identity               `json:"value"`
	Other       map[string]interface{} `json:"other"`
	Txid        string                 `json:"txid"`
	Vout        int                    `json:"vout,omitempty"`
	Address     string                 `json:"address"`
	Height      int64                  `json:"height,omitempty"`
	ExpiresIn   int64                  `json:"expires_in"`
	Expired     bool                   `json:"expired,omitempty"`
	Transferred bool                   `json:"transferred,omitempty"`
}

// Identity models fields that may be part of the JSON value of an identity.
type Identity struct {
	Name string `json:"name"`
	// String, array, or object with labels
	Email   interface{} `json:"email,omitempty"`
	Country string      `json:"country,omitempty"`
	// String or array
	Locality    interface{} `json:"locality,omitempty"`
	PhotoURL    string      `json:"photo_url,omitempty"`
	Description string      `json:"description,omitempty"`
	Hobby       []string    `json:"hobby,omitempty"`
	Birthday    string      `json:"birthday,omitempty"`
	Gender      string      `json:"gender,omitempty"`
	Weblog      string      `json:"weblog,omitempty"`
	// String, array, or object with labels
	Namecoin interface{} `json:"namecoin,omitempty"`
	// String, array, or object with labels
	Bitcoin    interface{} `json:"bitcoin,omitempty"`
	Bitmessage string      `json:"bitmessage,omitempty"`
	Xmpp       string      `json:"xmpp,omitempty"`
	// String or array of strings.
	// case insensitive, may contain spaces that are ignored.
	Otr interface{} `json:"otr,omitempty"`
	Gpg interface{} `json:"gpg,omitempty"`
	// String or array (addresses)
	Signer interface{} `json:"signer,omitempty"`
	// Custom fields
	Misc map[string]interface{} `json:",omitempty"`
}

// Used to prevent infinite recursion
type identity Identity

// Used to determine if a given key belongs in an IdentityShow's Misc field
var known_id_keys = map[string]bool{
	"name": true, "email": true, "country": true, "locality": true, "photo_url": true,
	"description": true, "hobby": true, "birthday": true, "gender": true,
	"weblog": true, "namecoin": true, "bitcoin": true, "bitmessage": true, "xmpp": true,
	"otr": true, "gpg": true, "signer": true,
}

// Satisfies the json.Unmarshaler interface
func (id *Identity) UnmarshalJSON(b []byte) error {
	idFields := make(map[string]interface{})
	miscFields := make(map[string]interface{})
	a := identity{}
	if err := json.Unmarshal(b, &a); err != nil {
		return err
	}
	// Misc fields
	err := json.Unmarshal(b, &idFields)
	if err != nil {
		return err
	}
	for k, v := range idFields {
		if !known_id_keys[k] {
			miscFields[k] = v
		}
	}
	a.Misc = miscFields
	*id = Identity(a)
	return nil
}
