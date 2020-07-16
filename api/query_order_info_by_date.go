package api

import (
	"fmt"
	"net/url"
)

type QueryOrderInfoByDate struct {
	sn    int
	date  string
	value url.Values
}

func NewQueryOrderInfoByDate(sn int, date string) *QueryOrderInfoByDate {
	object := &QueryOrderInfoByDate{}
	object.sn = sn
	object.date = date
	object.value = url.Values{}
	return object
}

func (receiver *QueryOrderInfoByDate) GetPostMsg(version string, key string) (msg string, err error) {
	if receiver.sn == 0 {
		return "", fmt.Errorf("required parameters cannot be null")
	}
	receiver.update()
	return UpdatePostValue(receiver.value, version, key), nil
}

func (receiver *QueryOrderInfoByDate) GetGetMsg(version string, key string) (msg string, err error) {
	if receiver.sn == 0 {
		return "", fmt.Errorf("required parameters cannot be null")
	}
	receiver.update()
	return UpdateGetValue(receiver.value, version, key), nil
}

func (receiver *QueryOrderInfoByDate) update() {
	receiver.value.Add("Action", "QueryOrderInfoByDate")
	receiver.value.Add("SN", fmt.Sprint(receiver.sn))
	receiver.value.Add("Date", receiver.date)
}
