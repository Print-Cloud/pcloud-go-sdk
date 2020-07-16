package api

import (
	"fmt"
	"net/url"
)

type QueryPrinterList struct {
	index  int
	number int
	value  url.Values
}

func NewQueryPrinterList(index int, number int) *QueryPrinterList {
	object := &QueryPrinterList{}
	object.index = index
	object.number = number
	object.value = url.Values{}
	return object
}

func (receiver *QueryPrinterList) GetPostMsg(version string, key string) (msg string, err error) {
	if receiver.index < 0 {
		return "", fmt.Errorf("required parameters cannot be null")
	}
	if receiver.number <= 0 || receiver.number >= 100 {
		return "", fmt.Errorf("required parameters cannot be null")
	}
	receiver.update()
	return UpdatePostValue(receiver.value, version, key), nil
}

func (receiver *QueryPrinterList) GetGetMsg(version string, key string) (msg string, err error) {
	if receiver.index < 0 {
		return "", fmt.Errorf("required parameters cannot be null")
	}
	if receiver.number <= 0 || receiver.number >= 100 {
		return "", fmt.Errorf("required parameters cannot be null")
	}
	receiver.update()
	return UpdateGetValue(receiver.value, version, key), nil
}

func (receiver *QueryPrinterList) update() {
	receiver.value.Add("Action", "QueryPrinterList")
	receiver.value.Add("Index", fmt.Sprint(receiver.index))
	receiver.value.Add("Number", fmt.Sprint(receiver.number))
}
