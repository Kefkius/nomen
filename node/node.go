package node

import (
	"encoding/json"
	"errors"
	"github.com/btcsuite/btcjson"
	"github.com/kefkius/nmcjson"
	"github.com/kefkius/nomen/config"
	"github.com/kefkius/nomen/models"
	"strings"
)

var (

	// rpcConf contains Namecoin RPC configuration info that the server uses.
	rpcConf config.Config

	// callId is used for RPC commands.
	callId int
)

// Init initializes rpcConf and callId
func Init(user, pass, server string) {
	rpcConf = config.New(user, pass, server)
	callId = 0
}

// SendRPC send an RPC command to the Namecoin node.
func SendRPC(cmd btcjson.Cmd) (btcjson.Reply, error) {
	return btcjson.RpcSend(rpcConf.User, rpcConf.Password, rpcConf.Server, cmd)
}

// GetBalance calls getbalance
func GetBalance() (float64, error) {
	callId += 1
	cmd, err := btcjson.NewGetBalanceCmd(callId)
	if err != nil {
		return 0.0, err
	}
	reply, err := SendRPC(cmd)
	if err != nil {
		return 0.0, err
	}
	if reply.Result != nil {
		if info, ok := reply.Result.(float64); ok {
			return info, nil
		}
	}
	return 0.0, errors.New("nil result")
}

// ShowName calls name_show.
// A prefix string may be passed.
func ShowName(name, prefix string) (nmcjson.NameShowResult, error) {
	fullName := name
	if prefix != "" {
		fullName = strings.Join([]string{prefix, name}, "")
	}
	callId += 1
	cmd, err := nmcjson.NewNameShowCmd(callId, fullName)
	if err != nil {
		return nmcjson.NameShowResult{}, err
	}
	reply, err := SendRPC(cmd)
	if err != nil {
		return nmcjson.NameShowResult{}, err
	}
	if reply.Result != nil {
		if info, ok := reply.Result.(nmcjson.NameShowResult); ok {
			return info, nil
		}
	}
	return nmcjson.NameShowResult{}, errors.New("nil result")
}

// ShowId calls name_show for an ID.
func ShowId(name string) (models.IdentityShow, error) {
	callId += 1
	fullName := strings.Join([]string{"id/", name}, "")
	cmd, err := nmcjson.NewNameShowCmd(callId, fullName)
	if err != nil {
		return models.IdentityShow{}, err
	}
	reply, err := SendRPC(cmd)
	if err != nil {
		return models.IdentityShow{}, err
	}
	if reply.Result == nil {
		return models.IdentityShow{}, errors.New("nil result")
	}
	res, ok := reply.Result.(nmcjson.NameShowResult)
	if !ok {
		return models.IdentityShow{}, errors.New("cannot parse result")
	}
	ident := &models.Identity{}
	val := res.Value
	err = json.Unmarshal([]byte(val), &ident)
	if err != nil {
		return models.IdentityShow{}, err
	}
	// Fill the Other field with identity's Misc value
	other := ident.Misc
	ident.Misc = nil
	id := models.IdentityShow{
		Name:      res.Name,
		Value:     *ident,
		Other:     other,
		Txid:      res.Txid,
		Vout:      res.Vout,
		Address:   res.Address,
		Height:    res.Height,
		ExpiresIn: res.ExpiresIn,
		Expired:   res.Expired,
	}
	return id, nil
}

// ShowDomain calls name_show for a domain.
func ShowDomain(name string) (models.DomainShow, error) {
	callId += 1
	fullName := strings.Join([]string{"d/", name}, "")
	cmd, err := nmcjson.NewNameShowCmd(callId, fullName)
	if err != nil {
		return models.DomainShow{}, err
	}
	reply, err := SendRPC(cmd)
	if err != nil {
		return models.DomainShow{}, err
	}
	if reply.Result == nil {
		return models.DomainShow{}, errors.New("nil result")
	}
	res, ok := reply.Result.(nmcjson.NameShowResult)
	if !ok {
		return models.DomainShow{}, errors.New("cannot parse result")
	}
	domain := &models.Domain{}
	val := res.Value
	err = json.Unmarshal([]byte(val), &domain)
	if err != nil {
		return models.DomainShow{}, err
	}
	// Fill the Other field with domain's Misc value
	other := domain.Misc
	domain.Misc = nil
	d := models.DomainShow{
		Name:      res.Name,
		Value:     *domain,
		Other:     other,
		Txid:      res.Txid,
		Vout:      res.Vout,
		Address:   res.Address,
		Height:    res.Height,
		ExpiresIn: res.ExpiresIn,
		Expired:   res.Expired,
	}
	return d, nil
}
