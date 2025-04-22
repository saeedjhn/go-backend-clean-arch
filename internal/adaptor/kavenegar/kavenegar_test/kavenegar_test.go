package kavenegar_test

import (
	"testing"

	"github.com/saeedjhn/go-backend-clean-arch/internal/adaptor/kavenegar"
)

func TestSendSingle_WithValidMessage_ReturnsMessageID(t *testing.T) {
	receptor := "09911228131"
	message := `سلام

این یک پیام آزمایشی است

لغو ۱۱`

	k := kavenegar.New(kavenegar.Config{
		Provider: "kavenegar",
		Sender:   "2000660110",
		APIKey:   "5A4E684A446A5945754E4A39507458357057343166446D417637494E534A5467686F6473655A772F476D633D",
	})
	recID, err := k.SendSingle(receptor, message)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("recID is: ", recID)
}
