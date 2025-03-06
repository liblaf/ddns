package notify

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/samber/oops"
	"github.com/spf13/viper"
)

type Bot struct {
	*gotgbot.Bot
	ChatID int64
}

func NewBot() (*Bot, error) {
	token := viper.GetString("telegram-token")
	chatId := viper.GetInt64("telegram-chat-id")
	raw, err := gotgbot.NewBot(token, &gotgbot.BotOpts{})
	if err != nil {
		return nil, oops.Wrap(err)
	}
	bot := &Bot{Bot: raw, ChatID: chatId}
	return bot, nil
}
