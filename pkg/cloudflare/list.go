package cloudflare

import (
	"context"

	cf "github.com/cloudflare/cloudflare-go/v5"
	"github.com/cloudflare/cloudflare-go/v5/dns"
)

type DNSRecord struct {
	ID string
}

func (c *Client) List(ctx context.Context, name string) ([]dns.RecordResponse, error) {
	records := []dns.RecordResponse{}
	iter := c.DNS.Records.ListAutoPaging(ctx, dns.RecordListParams{
		ZoneID: cf.F(c.ZoneID),
		Name: cf.F(dns.RecordListParamsName{
			Exact: cf.F(name),
		}),
	})
	for iter.Next() {
		records = append(records, iter.Current())
	}
	if err := iter.Err(); err != nil {
		return nil, iter.Err()
	}
	return records, nil
}
