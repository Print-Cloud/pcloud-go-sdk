package api

import (
	"fmt"
	"net/url"
)

type QueryOrderState struct {
	orderid string
	value   url.Values
}

func NewQueryOrderState(orderid string) *QueryOrderState {
	object := &QueryOrderState{}
	object.orderid = orderid
	object.value = url.Values{}
	return object
}

func (receiver *QueryOrderState) GetPostMsg(version string, key string) (msg string, err error) {
	if receiver.orderid == "" || len(receiver.orderid) < 20 {
		return "", fmt.Errorf("required parameters cannot be null")
	}
	receiver.update()
	return UpdatePostValue(receiver.value, version, key), nil
}

func (receiver *QueryOrderState) GetGetMsg(version string, key string) (msg string, err error) {
	if receiver.orderid == "" || len(receiver.orderid) < 20 {
		return "", fmt.Errorf("required parameters cannot be null")
	}
	receiver.update()
	return UpdateGetValue(receiver.value, version, key), nil
}

func (receiver *QueryOrderState) update() {
	receiver.value.Add("Action", "QueryOrderState")
	receiver.value.Add("Orderid", receiver.orderid)
}
