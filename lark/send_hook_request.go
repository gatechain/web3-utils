package lark

/*
  lark openapi 技术文档见： https://open.larksuite.com/document/client-docs/bot-v3/use-custom-bots-in-a-group
*/

import (
	"fmt"
	"net/http"
	"web3-utils/httputil"

	"github.com/pkg/errors"
)

const (
	hookUrl = "https://open.larksuite.com/open-apis/bot/v2/hook"
)

// secret 为空时不进行签名
func SendHookRequest(botID string, secret string, req *BotHookRequest) error {
	return SendHookRequestWithUrl(fmt.Sprintf("%s/%s", hookUrl, botID), secret, req)
}

// secret 为空时不进行签名
func SendHookRequestWithUrl(url string, secret string, req *BotHookRequest) error {
	if secret != "" {
		sign, timestamp, err := GenSignForNow(secret)
		if err != nil {
			return errors.WithMessage(err, "GenSignForNow err")
		}

		req.Timestamp = fmt.Sprintf("%d", timestamp)
		req.Sign = sign
	}
	resp := &BotHookResp{}
	err := httputil.SendHttpRequest(http.MethodPost, url, nil, req, resp)
	if err != nil {
		return errors.WithMessage(err, "SendHttpRequest err")
	}

	if resp.Code != 0 {
		return fmt.Errorf("bot hook response err, code=%d, msg=%s", resp.Code, resp.Message)
	}
	return nil
}
