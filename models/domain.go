package models

import (
	"encoding/json"
)

// DomainShow models fields in a name_show result
type DomainShow struct {
	Name        string                 `json:"name"`
	Value       Domain                 `json:"value"`
	Other       map[string]interface{} `json:"other"`
	Txid        string                 `json:"txid"`
	Vout        int                    `json:"vout,omitempty"`
	Address     string                 `json:"address"`
	Height      int64                  `json:"height,omitempty"`
	ExpiresIn   int64                  `json:"expires_in"`
	Expired     bool                   `json:"expired,omitempty"`
	Transferred bool                   `json:"transferred,omitempty"`
}

// Domain models fields that may be part of the JSON value of a domain.
type Domain struct {
	Service interface{} `json:"service,omitempty"`
	// String or array
	Ip interface{} `json:"ip"`
	// String or array
	Ip6       interface{} `json:"ip6,omitempty"`
	Tor       string      `json:"tor,omitempty"`
	I2p       interface{} `json:"i2p,omitempty"`
	Freenet   string      `json:"freenet,omitempty"`
	Alias     string      `json:"alias,omitempty"`
	Translate string      `json:"translate,omitempty"`
	Email     string      `json:"email,omitempty"`
	Loc       string      `json:"loc,omitempty"`
	// Registrant info format?
	Info        interface{} `json:"info,omitempty"`
	Ns          interface{} `json:"ns,omitempty"`
	Delegate    interface{} `json:"delegate,omitempty"`
	Import      interface{} `json:"import,omitempty"`
	Map         interface{} `json:"map,omitempty"`
	Fingerprint interface{} `json:"fingerprint,omitempty"`
	Tls         interface{} `json:"tls,omitempty"`
	Ds          interface{} `json:"ds,omitempty"`
	// Custom fields
	Misc map[string]interface{} `json:",omitempty"`
}

// Used to prevent infinite recursion
type domain Domain

// Used to determine if a given key belongs in a DomainShow's Misc field
var known_domain_keys = map[string]bool{
	"service": true, "ip": true, "ip6": true, "tor": true, "i2p": true, "freenet": true,
	"alias": true, "translate": true, "email": true, "loc": true, "info": true, "ns": true,
	"delegate": true, "import": true, "map": true, "fingerprint": true, "tls": true, "ds": true,
}

// Satisfies the json.Unmarshaler interface
func (d *Domain) UnmarshalJSON(b []byte) error {
	dFields := make(map[string]interface{})
	miscFields := make(map[string]interface{})
	a := domain{}
	if err := json.Unmarshal(b, &a); err != nil {
		return err
	}
	// Misc fields
	err := json.Unmarshal(b, &dFields)
	if err != nil {
		return err
	}
	for k, v := range dFields {
		if !known_domain_keys[k] {
			miscFields[k] = v
		}
	}
	a.Misc = miscFields
	*d = Domain(a)
	return nil
}
