package httputil

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

func SendHttpRequest(method, url string, headers map[string]string, payload, result any) error {
	var body io.Reader
	if payload != nil {
		data, err := json.Marshal(payload)
		if err != nil {
			return errors.WithMessage(err, "marshal payload err")
		}
		body = bytes.NewBuffer(data)
	}

	req, _ := http.NewRequest(method, url, body)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return errors.WithMessage(err, "do request err")
	}
	defer resp.Body.Close()

	if result != nil {
		err := json.NewDecoder(resp.Body).Decode(result)
		if err != nil {
			return errors.WithMessage(err, "decode response body err")
		}
	}
	return nil
}
