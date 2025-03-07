package cloudflare

import (
	"context"
	"net/netip"

	cf "github.com/cloudflare/cloudflare-go/v4"
	"github.com/cloudflare/cloudflare-go/v4/dns"
	"github.com/rs/zerolog/log"
	"github.com/samber/oops"
)

func (c *Client) Update(
	ctx context.Context,
	comment string,
	name string,
	deletes []dns.RecordResponse,
	posts []netip.Addr,
) (*dns.RecordBatchResponse, error) {
	if len(deletes) == 0 && len(posts) == 0 {
		return nil, nil
	}
	params := dns.RecordBatchParams{
		ZoneID: cf.F(c.ZoneID),
	}
	if len(deletes) > 0 {
		deletesParams := []dns.RecordBatchParamsDelete{}
		for _, record := range deletes {
			deletesParams = append(deletesParams, dns.RecordBatchParamsDelete{
				ID: cf.F(record.ID),
			})
		}
		params.Deletes = cf.F(deletesParams)
	}
	if len(posts) > 0 {
		postsParams := []dns.RecordUnionParam{}
		for _, ip := range posts {
			if ip.Is4() {
				postsParams = append(postsParams, dns.ARecordParam{
					Comment: cf.F(comment),
					Content: cf.F(ip.String()),
					Name:    cf.F(name),
					Proxied: cf.F(false),
					TTL:     cf.F(dns.TTL(60)),
					Type:    cf.F(dns.ARecordTypeA),
				})
			} else if ip.Is6() {
				postsParams = append(postsParams, dns.AAAARecordParam{
					Comment: cf.F(comment),
					Content: cf.F(ip.String()),
					Name:    cf.F(name),
					Proxied: cf.F(false),
					TTL:     cf.F(dns.TTL(60)),
					Type:    cf.F(dns.AAAARecordTypeAAAA),
				})
			} else {
				log.Warn().IPAddr("ip", ip.AsSlice()).Send()
			}
		}
		params.Posts = cf.F(postsParams)
	}
	resp, err := c.DNS.Records.Batch(ctx, params)
	if err != nil {
		return nil, oops.Wrap(err)
	}
	log.Debug().Any("response", resp).Send()
	for _, record := range resp.Deletes {
		log.Info().Msgf("Delete %s => %s", GetLabel(record), record.Content)
	}
	for _, record := range resp.Posts {
		log.Info().Msgf("Post %s => %v", GetLabel(record), record.Content)
	}
	return resp, nil
}
