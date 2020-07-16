package api

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type PrinterInfoItem struct {
	SN     int    `json:"SN"`
	Remark string `json:"Remark"`
}

type PrinterAdd struct {
	list  []*PrinterInfoItem
	value url.Values
}

func NewPrinterAdd(list []*PrinterInfoItem) *PrinterAdd {
	object := &PrinterAdd{}
	object.list = list
	object.value = url.Values{}
	return object
}

func (receiver *PrinterAdd) GetPostMsg(version string, key string) (msg string, err error) {
	if len(receiver.list) == 0 {
		return "", fmt.Errorf("required parameters cannot be null")
	}
	_err := receiver.update()
	if _err != nil {
		return "", _err
	}
	return UpdatePostValue(receiver.value, version, key), nil
}

func (receiver *PrinterAdd) GetGetMsg(version string, key string) (msg string, err error) {
	if len(receiver.list) == 0 {
		return "", fmt.Errorf("required parameters cannot be null")
	}
	_err := receiver.update()
	if _err != nil {
		return "", _err
	}
	return UpdateGetValue(receiver.value, version, key), nil
}

func (receiver *PrinterAdd) update() error {
	_bytes, _err := json.Marshal(receiver.list)
	if _err != nil {
		return _err
	}
	receiver.value.Add("Action", "PrinterAdd")
	receiver.value.Add("List", string(_bytes))
	return nil
}
