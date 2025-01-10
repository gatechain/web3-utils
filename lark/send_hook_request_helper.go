package lark

/*
  lark openapi 技术文档见： https://open.larksuite.com/document/client-docs/bot-v3/use-custom-bots-in-a-group
*/

import (
	"fmt"
)

// secret 为空时不进行签名
func SendHookText(botID, secret, message string) error {
	return SendHookTextWithUrl(fmt.Sprintf("%s/%s", hookUrl, botID), secret, message)
}

// secret 为空时不进行签名
func SendHookTextWithUrl(url, secret, message string) error {
	req := &BotHookRequest{
		MsgType: MsgType_Text,
		Content: &Content{
			Text: message,
		},
	}
	return SendHookRequestWithUrl(url, secret, req)
}

// secret 为空时不进行签名
func SendHookRichTextMsg(botID, secret, title string, firstLine RichTextLine, moreLines ...RichTextLine) error {
	return SendHookRichTextMsgWithUrl(fmt.Sprintf("%s/%s", hookUrl, botID), secret, title, firstLine, moreLines...)
}

// secret 为空时不进行签名
func SendHookRichTextMsgWithUrl(url, secret, title string, firstLine RichTextLine, moreLines ...RichTextLine) error {
	req := &BotHookRequest{
		MsgType: MsgType_Post,
		Content: &Content{
			Post: &Post{
				ZhCn: &RichTextContent{
					Title:   title,
					Content: append([]RichTextLine{firstLine}, moreLines...),
				},
			},
		},
	}
	return SendHookRequestWithUrl(url, secret, req)
}
