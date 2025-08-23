package notify

import (
	"fmt"

	"github.com/cloudflare/cloudflare-go/v5/dns"
	"github.com/liblaf/ddns/pkg/cloudflare"
)

func PrettyRecord(emoji string, record dns.RecordResponse) string {
	return fmt.Sprintf(
		"%s <a href=\"%s\">%s</a> => <code>%s</code>",
		emoji,
		record.Name,
		cloudflare.GetLabel(record),
		record.Content,
	)
}
