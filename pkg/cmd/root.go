package cmd

import (
	"github.com/liblaf/ddns/pkg/logging"
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
				viper.SetConfigType("yaml")
			}
			logging.Init()
			if err := viper.ReadInConfig(); err != nil {
				return oops.Wrap(err)
			}
			return nil
		},
	}
	cmd.AddCommand(Delete(), List(), Update())
	cmd.PersistentFlags().StringP("config", "c", "", "")
	cmd.PersistentFlags().StringP("domain", "d", "", "")
	cmd.PersistentFlags().StringP("telegram-chat-id", "C", "", "")
	cmd.PersistentFlags().StringP("telegram-token", "T", "", "")
	cmd.PersistentFlags().StringP("token", "t", "", "")
	cmd.PersistentFlags().StringP("zone-id", "z", "", "")
	viper.BindPFlag("domain", cmd.Flags().Lookup("domain"))
	viper.BindPFlag("telegram-chat-id", cmd.Flags().Lookup("telegram-chat-id"))
	viper.BindPFlag("telegram-token", cmd.Flags().Lookup("telegram-token"))
	viper.BindPFlag("token", cmd.Flags().Lookup("token"))
	viper.BindPFlag("zone-id", cmd.Flags().Lookup("zone-id"))
	return cmd
}
