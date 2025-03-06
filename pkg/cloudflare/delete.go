package cloudflare

import (
	"context"

	cf "github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/dns"
	"github.com/rs/zerolog/log"
	"github.com/samber/oops"
)

func (c *Client) Delete(ctx context.Context, records []dns.RecordResponse) error {
	if len(records) == 0 {
		return nil
	}
	deletes := make([]dns.RecordBatchParamsDelete, 0, len(records))
	for _, record := range records {
		deletes = append(deletes, dns.RecordBatchParamsDelete{
			ID: cf.F(record.ID),
		})
	}
	resp, err := c.DNS.Records.Batch(ctx, dns.RecordBatchParams{
		ZoneID:  cf.F(c.ZoneID),
		Deletes: cf.F(deletes),
	})
	if err != nil {
		return oops.Wrap(err)
	}
	log.Debug().Any("response", resp).Send()
	for _, record := range resp.Deletes {
		log.Info().Msgf("Delete %s => %s", GetLabel(record), record.Content)
	}
	return nil
}
