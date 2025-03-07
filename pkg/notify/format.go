package notify

import (
	"fmt"

	"github.com/cloudflare/cloudflare-go/v4/dns"
)

func GetLabel(record dns.RecordResponse) string {
	if record.Comment != "" {
		return record.Comment
	}
	return record.Name
}

func PrettyRecord(emoji string, record dns.RecordResponse) string {
	return fmt.Sprintf(
		"%s <a href=\"%s\">%s</a> => <code>%s</code>",
		emoji,
		record.Name,
		GetLabel(record),
		record.Content,
	)
}
