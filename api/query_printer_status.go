package api

import (
	"fmt"
	"net/url"
)

type QueryPrinterStatus struct {
	sn    int
	value url.Values
}

func NewQueryPrinterStatus(sn int) *QueryPrinterStatus {
	object := &QueryPrinterStatus{}
	object.sn = sn
	object.value = url.Values{}
	return object
}

func (receiver *QueryPrinterStatus) GetPostMsg(version string, key string) (msg string, err error) {
	if receiver.sn == 0 {
		return "", fmt.Errorf("required parameters cannot be null")
	}
	receiver.update()
	return UpdatePostValue(receiver.value, version, key), nil
}

func (receiver *QueryPrinterStatus) GetGetMsg(version string, key string) (msg string, err error) {
	if receiver.sn == 0 {
		return "", fmt.Errorf("required parameters cannot be null")
	}
	receiver.update()
	return UpdateGetValue(receiver.value, version, key), nil
}

func (receiver *QueryPrinterStatus) update() {
	receiver.value.Add("Action", "QueryPrinterStatus")
	receiver.value.Add("SN", fmt.Sprint(receiver.sn))
}
