package node

import (
	"encoding/json"
	"errors"
	"github.com/kefkius/nmcjson"
	"github.com/kefkius/nomen/models"
	"strings"
)

// ListNames calls name_list
func ListNames() ([]nmcjson.NameListResult, error) {
	callId += 1
	cmd, err := nmcjson.NewNameListCmd(callId)
	if err != nil {
		return nil, err
	}
	reply, err := SendRPC(cmd)
	if err != nil {
		return nil, err
	}
	if reply.Result != nil {
		if info, ok := reply.Result.([]nmcjson.NameListResult); ok {
			return info, nil
		}
	}
	return nil, errors.New("nil result")
}

// ListIds calls name_list and filters identities
func ListIds() ([]models.IdentityShow, error) {
	allNames, err := ListNames()
	if err != nil {
		return nil, err
	}
	identities := []models.IdentityShow{}
	for _, v := range allNames {
		if strings.HasPrefix(v.Name, "id/") {
			ident := &models.Identity{}
			err = json.Unmarshal([]byte(v.Value), &ident)
			if err != nil {
				continue
			}
			id := models.IdentityShow{
				Name:      v.Name,
				Value:     *ident,
				Txid:      v.Txid,
				Vout:      v.Vout,
				Address:   v.Address,
				Height:    v.Height,
				ExpiresIn: v.ExpiresIn,
				Expired:   v.Expired,
			}
			identities = append(identities, id)
		}
	}
	return identities, nil
}

// ListDomains calls name_list and filters domains
func ListDomains() ([]models.DomainShow, error) {
	allNames, err := ListNames()
	if err != nil {
		return nil, err
	}
	domains := []models.DomainShow{}
	for _, v := range allNames {
		if strings.HasPrefix(v.Name, "d/") {
			domain := &models.Domain{}
			err = json.Unmarshal([]byte(v.Value), &domain)
			if err != nil {
				continue
			}
			d := models.DomainShow{
				Name:      v.Name,
				Value:     *domain,
				Txid:      v.Txid,
				Vout:      v.Vout,
				Address:   v.Address,
				Height:    v.Height,
				ExpiresIn: v.ExpiresIn,
				Expired:   v.Expired,
			}
			domains = append(domains, d)
		}
	}
	return domains, nil
}

// ListExpiredNames shows all expired names
func ListExpiredNames() ([]string, error) {
	allNames, err := ListNames()
	if err != nil {
		return nil, err
	}
	expiredNames := []string{}
	for _, v := range allNames {
		if v.Expired == true {
			expiredNames = append(expiredNames, v.Name)
		}
	}
	return expiredNames, nil
}
