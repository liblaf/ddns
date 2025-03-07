package cmd

import (
	"github.com/liblaf/ddns/pkg/cloudflare"
	"github.com/liblaf/ddns/pkg/config"
	"github.com/liblaf/ddns/pkg/notify"
	"github.com/samber/oops"
	"github.com/spf13/cobra"
)

func Delete() *cobra.Command {
	cmd := &cobra.Command{
		Use: "delete",
		RunE: func(cmd *cobra.Command, args []string) error {
			domain, err := config.Domain()
			if err != nil {
				return oops.Wrap(err)
			}
			client := cloudflare.NewClient()
			bot, err := notify.NewBot()
			if err != nil {
				return oops.Wrap(err)
			}
			records, err := client.List(cmd.Context(), domain)
			if err != nil {
				return oops.Wrap(err)
			}
			if len(records) == 0 {
				return nil
			}
			resp, err := client.Delete(cmd.Context(), records)
			if err != nil {
				return err
			}
			err = bot.Delete(resp.Deletes)
			if err != nil {
				return oops.Wrap(err)
			}
			return nil
		},
	}
	return cmd
}
