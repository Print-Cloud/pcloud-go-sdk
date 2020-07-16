package pcloud_go_sdk

import (
	"github.com/print-cloud/pcloud-go-sdk/api"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const Version = "1.0"

type Client struct {
	ApiUrl  string
	Key     string
	timeout time.Duration
}

func NewClient(url string, key string) *Client {
	return &Client{
		ApiUrl: url,
		Key:    key,
	}
}

func (s *Client) GetSend(api api.Api) (data string, err error) {
	client := http.Client{}
	if s.timeout != 0 {
		client.Timeout = s.timeout
	}
	_msg, _err := api.GetGetMsg(Version, s.Key)
	if _err != nil {
		return "", _err
	}
	res, _err := client.Get(s.ApiUrl + "?" + _msg)
	if _err != nil {
		return "", _err
	}
	defer res.Body.Close()

	_data, _err := ioutil.ReadAll(res.Body)
	if _err != nil {
		return "", _err
	}
	return string(_data), nil
}

func (s *Client) PostSend(api api.Api) (data string, err error) {
	client := http.Client{}
	if s.timeout != 0 {
		client.Timeout = s.timeout
	}
	_msg, _err := api.GetPostMsg(Version, s.Key)
	if _err != nil {
		return "", _err
	}
	res, _err := client.Post(s.ApiUrl, "application/x-www-form-urlencoded",
		strings.NewReader(_msg))
	if _err != nil {
		return "", _err
	}
	defer res.Body.Close()

	_data, _err := ioutil.ReadAll(res.Body)
	if _err != nil {
		return "", _err
	}

	return string(_data), nil
}

func (s *Client) WithTimeout(d time.Duration) {
	s.timeout = d
}
