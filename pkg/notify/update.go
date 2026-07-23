package notify

import (
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/cloudflare/cloudflare-go/v7/dns"
	"github.com/rs/zerolog/log"
	"github.com/samber/oops"
)

func (b *Bot) Update(keeps []dns.RecordResponse, deletes []dns.RecordResponse, posts []dns.RecordResponse) error {
	text := ""
	for _, record := range keeps {
		text += PrettyRecord("🔵", record) + "\n"
	}
	for _, record := range deletes {
		text += PrettyRecord("🔴", record) + "\n"
	}
	for _, record := range posts {
		text += PrettyRecord("🟢", record) + "\n"
	}
	text = strings.TrimSpace(text)
	message, err := b.SendMessage(b.ChatID, text, &gotgbot.SendMessageOpts{ParseMode: "HTML"})
	log.Debug().Any("message", message).Send()
	if err != nil {
		return oops.Wrap(err)
	}
	return nil
}
