package node

import (
	"errors"
	"github.com/kefkius/nmcjson"
)

// FilterNames calls name_filter.
func FilterNames(regexp string, numResults int) ([]string, error) {
	callId += 1
	if numResults == 0 {
		numResults = 20
	}
	cmd, err := nmcjson.NewNameFilterCmd(callId, regexp, 0, 0, numResults)
	if err != nil {
		return nil, err
	}
	reply, err := SendRPC(cmd)
	if err != nil {
		return nil, err
	}
	if reply.Result == nil {
		return nil, errors.New("nil result")
	}
	res, ok := reply.Result.([]nmcjson.NameFilterResult)
	if !ok {
		return nil, errors.New("cannot parse result")
	}
	nameList := []string{}
	for _, v := range res {
		nameList = append(nameList, v.Name)
	}
	return nameList, nil
}
