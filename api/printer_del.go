package api

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type PrinterDel struct {
	list  []*PrinterInfoItem
	value url.Values
}

func NewPrinterDel(list []*PrinterInfoItem) *PrinterDel {
	object := &PrinterDel{}
	object.list = list
	object.value = url.Values{}
	return object
}

func (receiver *PrinterDel) GetPostMsg(version string, key string) (msg string, err error) {
	if len(receiver.list) == 0 {
		return "", fmt.Errorf("required parameters cannot be null")
	}
	_err := receiver.update()
	if _err != nil {
		return "", _err
	}
	return UpdatePostValue(receiver.value, version, key), nil
}

func (receiver *PrinterDel) GetGetMsg(version string, key string) (msg string, err error) {
	if len(receiver.list) == 0 {
		return "", fmt.Errorf("required parameters cannot be null")
	}
	_err := receiver.update()
	if _err != nil {
		return "", _err
	}
	return UpdateGetValue(receiver.value, version, key), nil
}

func (receiver *PrinterDel) update() error {
	_bytes, _err := json.Marshal(receiver.list)
	if _err != nil {
		return _err
	}
	receiver.value.Add("Action", "PrinterDel")
	receiver.value.Add("List", string(_bytes))
	return nil
}
