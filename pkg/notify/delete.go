package notify

import (
	"fmt"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/cloudflare/cloudflare-go/v4/dns"
	"github.com/liblaf/ddns/pkg/cloudflare"
	"github.com/rs/zerolog/log"
	"github.com/samber/oops"
)

func (b *Bot) Delete(records []dns.RecordResponse) error {
	text := ""
	for _, record := range records {
		text += fmt.Sprintf("🔴 <code>%s</code> => <code>%s</code>\n", cloudflare.GetLabel(record), record.Content)
	}
	message, err := b.SendMessage(5596425538, text, &gotgbot.SendMessageOpts{ParseMode: "HTML"})
	log.Debug().Any("message", message).Send()
	if err != nil {
		return oops.Wrap(err)
	}
	return nil
}
