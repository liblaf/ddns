package cmd

import (
	"github.com/liblaf/ddns/pkg/logging"
	"github.com/rs/zerolog"
	"github.com/samber/oops"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Root() *cobra.Command {
	cmd := &cobra.Command{
		Use: "ddns",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			config, err := cmd.Flags().GetString("config")
			if err != nil {
				return oops.Wrap(err)
			}
			if config != "" {
				viper.SetConfigFile(config)
			} else {
				viper.AddConfigPath(".")
				viper.AddConfigPath("$HOME/.config/liblaf/ddns")
				viper.AddConfigPath("/etc/liblaf/ddns")
				viper.SetConfigName("config")
				viper.SetConfigType("toml")
			}
			verbose, err := cmd.Flags().GetCount("verbose")
			if err != nil {
				return oops.Wrap(err)
			}
			quiet, err := cmd.Flags().GetCount("quiet")
			if err != nil {
				return oops.Wrap(err)
			}
			logging.Init(zerolog.Level(quiet - verbose))
			if err := viper.ReadInConfig(); err != nil {
				return oops.Wrap(err)
			}
			return nil
		},
	}
	cmd.AddCommand(Delete(), List(), Update())
	cmd.PersistentFlags().CountP("quiet", "q", "")
	cmd.PersistentFlags().CountP("verbose", "v", "")
	cmd.PersistentFlags().String("telegram-bot-token", "", "")
	cmd.PersistentFlags().String("telegram-chat-id", "", "")
	cmd.PersistentFlags().StringP("config", "c", "", "(default [\"config.toml\", \"~/.config/liblaf/ddns/config.toml\", \"/etc/liblaf/ddns/config.toml\"])")
	cmd.PersistentFlags().StringP("domain", "d", "", "(default \"${HOSTNAME}.ddns.liblaf.me\")")
	cmd.PersistentFlags().StringP("token", "t", "", "")
	cmd.PersistentFlags().StringP("zone-id", "z", "", "")
	viper.BindEnv("domain", "DOMAIN")
	viper.BindEnv("telegram-bot-token", "TELEGRAM_BOT_TOKEN")
	viper.BindEnv("telegram-chat-id", "TELEGRAM_CHAT_ID")
	viper.BindEnv("token", "CLOUDFLARE_API_TOKEN")
	viper.BindEnv("zone-id", "CLOUDFLARE_ZONE_ID")
	viper.BindPFlag("domain", cmd.Flags().Lookup("domain"))
	viper.BindPFlag("telegram-bot-token", cmd.Flags().Lookup("telegram-bot-token"))
	viper.BindPFlag("telegram-chat-id", cmd.Flags().Lookup("telegram-chat-id"))
	viper.BindPFlag("token", cmd.Flags().Lookup("token"))
	viper.BindPFlag("zone-id", cmd.Flags().Lookup("zone-id"))
	return cmd
}
