package lark

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"time"
)

func GenSign(secret string, timestamp int64) (string, error) {
	//timestamp + key 做sha256, 再进行base64 encode
	stringToSign := fmt.Sprintf("%v", timestamp) + "\n" + secret

	var data []byte
	h := hmac.New(sha256.New, []byte(stringToSign))
	_, err := h.Write(data)
	if err != nil {
		return "", err
	}

	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return signature, nil
}

func GenSignForNow(secret string) (sign string, timestamp int64, err error) {
	timestamp = time.Now().Unix()
	sign, err = GenSign(secret, timestamp)
	if err != nil {
		return "", 0, err
	}
	return sign, timestamp, nil
}
