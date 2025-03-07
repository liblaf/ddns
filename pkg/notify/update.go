package notify

import (
	"fmt"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/cloudflare/cloudflare-go/v4/dns"
	"github.com/liblaf/ddns/pkg/cloudflare"
	"github.com/rs/zerolog/log"
	"github.com/samber/oops"
)

func (b *Bot) Update(keeps []dns.RecordResponse, deletes []dns.RecordResponse, posts []dns.RecordResponse) error {
	text := ""
	for _, record := range keeps {
		text += fmt.Sprintf("🔵 <code>%s</code> => <code>%s</code>\n", cloudflare.GetLabel(record), record.Content)
	}
	for _, record := range deletes {
		text += fmt.Sprintf("🔴 <code>%s</code> => <code>%s</code>\n", cloudflare.GetLabel(record), record.Content)
	}
	for _, record := range posts {
		text += fmt.Sprintf("🟢 <code>%s</code> => <code>%s</code>\n", cloudflare.GetLabel(record), record.Content)
	}
	text = strings.TrimSpace(text)
	message, err := b.SendMessage(b.ChatID, text, &gotgbot.SendMessageOpts{ParseMode: "HTML"})
	log.Debug().Any("message", message).Send()
	if err != nil {
		return oops.Wrap(err)
	}
	return nil
}
