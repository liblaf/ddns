package notify

import (
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/cloudflare/cloudflare-go/v4/dns"
	"github.com/rs/zerolog/log"
	"github.com/samber/oops"
)

func (b *Bot) Delete(records []dns.RecordResponse) error {
	text := ""
	for _, record := range records {
		text += PrettyRecord("🔴", record) + "\n"
	}
	text = strings.TrimSpace(text)
	message, err := b.SendMessage(b.ChatID, text, &gotgbot.SendMessageOpts{ParseMode: "HTML"})
	log.Debug().Any("message", message).Send()
	if err != nil {
		return oops.Wrap(err)
	}
	return nil
}
