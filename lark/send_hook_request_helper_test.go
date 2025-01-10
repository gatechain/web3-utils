package lark

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	botID  = "e22fa968-ceed-453d-ab41-3ba12d5c72df"
	secret = "8TBQ0xEQjIwCM6fUbbH4qc"

	// user-open-id
	hogan   = "ou_93f07a888ebb199d5b079e05d396f14c"
	wesley  = "ou_0d59602926329d76e0cdf938221b3b6e"
	athos   = "ou_5fc360ea18fff8e5128e214147709b98"
	clement = "ou_7a32839f52df9aebebd1bc0c6b59c557"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestSendHookText(t *testing.T) {
	err := SendHookText(botID, "", "this is a message from web3-utils/lark")
	assert.Error(t, err)

	err = SendHookText(botID, secret, "this is a message from web3-utils/lark")
	assert.NoError(t, err)
}

func TestSendHookRichTextMsg(t *testing.T) {
	SendHookRichTextMsg(botID, secret, "Happy Friday Night",
		RichTextLine{
			RichTextItemText("恭喜发财:"),
			RichTextItemAt(hogan),
			RichTextItemHref("交易就上 gate.io", "https://www.gate.io"),
		},
		RichTextLine{
			RichTextItemText("升职加薪:"),
			RichTextItemAt(athos),
			RichTextItemAt(hogan),
			RichTextItemAt(wesley),
			RichTextItemAt(clement),
		},
		RichTextLine{
			RichTextItemText("步步高:"),
			RichTextItemAt(athos),
			RichTextItemAt(hogan),
			RichTextItemAt(wesley),
			RichTextItemAt(clement),
		},
	)
}
