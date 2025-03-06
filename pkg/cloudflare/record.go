package cloudflare

import "github.com/cloudflare/cloudflare-go/v4/dns"

func GetLabel(record dns.RecordResponse) string {
	if record.Comment != "" {
		return record.Comment
	}
	return record.Name
}
