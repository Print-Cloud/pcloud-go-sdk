package api

import (
	"fmt"
	"net/url"
)

type DelPrinterSqs struct {
	sn    int
	value url.Values
}

func NewDelPrinterSqs(sn int) *DelPrinterSqs {
	object := &DelPrinterSqs{}
	object.sn = sn
	object.value = url.Values{}
	return object
}

func (receiver *DelPrinterSqs) GetPostMsg(version string, key string) (msg string, err error) {
	if receiver.sn == 0 {
		return "", fmt.Errorf("required parameters cannot be null")
	}
	_ = receiver.update()

	return UpdatePostValue(receiver.value, version, key), nil
}

func (receiver *DelPrinterSqs) GetGetMsg(version string, key string) (msg string, err error) {
	if receiver.sn == 0 {
		return "", fmt.Errorf("required parameters cannot be null")
	}

	_ = receiver.update()

	return UpdateGetValue(receiver.value, version, key), nil
}

func (receiver *DelPrinterSqs) update() error {
	receiver.value.Add("Action", "DelPrinterSqs")
	receiver.value.Add("SN", fmt.Sprint(receiver.sn))
	return nil
}
