package api

import (
	"fmt"
	"net/url"
)

type PrintMsg struct {
	sn      int
	content string
	value   url.Values
}

func NewPrintMsg(sn int, content string) *PrintMsg {
	object := &PrintMsg{}
	object.sn = sn
	object.content = content
	object.value = url.Values{}
	return object
}

func (receiver *PrintMsg) WithTimes(times int) {
	receiver.value.Add("Times", fmt.Sprint(times))
}

func (receiver *PrintMsg) WithLanguage(language string) {
	receiver.value.Add("Language", language)
}

func (receiver *PrintMsg) GetPostMsg(version string, key string) (msg string, err error) {
	if receiver.sn == 0 {
		return "", fmt.Errorf("required parameters cannot be null")
	}
	if receiver.content == "" {
		return "", fmt.Errorf("required parameters cannot be null")
	}

	_ = receiver.update()

	return UpdatePostValue(receiver.value, version, key), nil
}

func (receiver *PrintMsg) GetGetMsg(version string, key string) (msg string, err error) {
	if receiver.sn == 0 {
		return "", fmt.Errorf("required parameters cannot be null")
	}
	if receiver.content == "" {
		return "", fmt.Errorf("required parameters cannot be null")
	}

	_ = receiver.update()

	return UpdateGetValue(receiver.value, version, key), nil
}

func (receiver *PrintMsg) update() error {
	receiver.value.Add("Action", "PrintMsg")
	receiver.value.Add("SN", fmt.Sprint(receiver.sn))
	receiver.value.Add("Content", receiver.content)
	return nil
}
