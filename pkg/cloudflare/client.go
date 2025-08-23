package cloudflare

import (
	cf "github.com/cloudflare/cloudflare-go/v5"
	"github.com/cloudflare/cloudflare-go/v5/option"
	"github.com/spf13/viper"
)

type Client struct {
	*cf.Client
	ZoneID string
}

func NewClient() Client {
	token := viper.GetString("token")
	client := Client{
		Client: cf.NewClient(
			option.WithAPIToken(token),
		),
		ZoneID: viper.GetString("zone-id"),
	}
	return client
}
