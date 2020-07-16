package api

import (
	"fmt"
	"net/url"
)

type PrinterEdit struct {
	sn     int
	remark string
	value  url.Values
}

func NewPrinterEdit(sn int, remark string) *PrinterEdit {
	object := &PrinterEdit{}
	object.sn = sn
	object.remark = remark
	object.value = url.Values{}
	return object
}

func (receiver *PrinterEdit) GetPostMsg(version string, key string) (msg string, err error) {
	if receiver.sn == 0 {
		return "", fmt.Errorf("required parameters cannot be null")
	}
	receiver.update()
	return UpdatePostValue(receiver.value, version, key), nil
}

func (receiver *PrinterEdit) GetGetMsg(version string, key string) (msg string, err error) {
	if receiver.sn == 0 {
		return "", fmt.Errorf("required parameters cannot be null")
	}
	receiver.update()
	return UpdateGetValue(receiver.value, version, key), nil
}

func (receiver *PrinterEdit) update() {
	receiver.value.Add("Action", "PrinterEdit")
	receiver.value.Add("SN", fmt.Sprint(receiver.sn))
	receiver.value.Add("Remark", receiver.remark)
}
