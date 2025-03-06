package cmd

import (
	"github.com/liblaf/ddns/pkg/cloudflare"
	"github.com/liblaf/ddns/pkg/config"
	"github.com/rs/zerolog/log"
	"github.com/samber/oops"
	"github.com/spf13/cobra"
)

func List() *cobra.Command {
	cmd := &cobra.Command{
		Use: "list",
		RunE: func(cmd *cobra.Command, args []string) error {
			domain, err := config.Domain()
			if err != nil {
				return oops.Wrap(err)
			}
			client := cloudflare.NewClient()
			records, err := client.List(cmd.Context(), domain)
			if err != nil {
				return err
			}
			for _, record := range records {
				log.Info().Msgf("%s => %s", cloudflare.GetLabel(record), record.Content)
			}
			return nil
		},
	}
	return cmd
}
