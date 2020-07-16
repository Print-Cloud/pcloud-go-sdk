package api

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/url"
	"time"
)

func UpdatePostValue(value url.Values, version string, key string) (encode string) {
	value.Add("Version", version)
	value.Add("Key", key)
	value.Add("Timestamp", fmt.Sprint(time.Now().Unix()))
	value.Add("Signature", PostSignature(value.Encode()))
	value.Del("Key")
	return value.Encode()
}

func UpdateGetValue(value url.Values, version string, key string) (encode string) {
	value.Add("Version", version)
	value.Add("Key", key)
	value.Add("Timestamp", fmt.Sprint(time.Now().Unix()))
	value.Add("Signature", GetSignature(value.Encode()))
	value.Del("Key")
	return value.Encode()
}

func PostSignature(encode string) string {
	_source := "POST?" + encode
	sum := sha256.Sum256([]byte(_source))
	encodeString := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%x", sum)))
	return encodeString
}

func GetSignature(encode string) string {
	_source := "GET?" + encode
	sum := sha256.Sum256([]byte(_source))
	encodeString := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%x", sum)))
	return encodeString
}
