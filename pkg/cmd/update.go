package cmd

import (
	"net/netip"

	"github.com/cloudflare/cloudflare-go/v4/dns"
	"github.com/emirpasic/gods/sets/hashset"
	"github.com/liblaf/ddns/pkg/cloudflare"
	"github.com/liblaf/ddns/pkg/config"
	"github.com/liblaf/ddns/pkg/ip"
	"github.com/liblaf/ddns/pkg/notify"
	"github.com/samber/oops"
	"github.com/spf13/cobra"
)

func Update() *cobra.Command {
	cmd := &cobra.Command{
		Use: "update",
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
			ips, err := ip.GlobalIPs()
			if err != nil {
				return oops.Wrap(err)
			}
			ipSet := StringSetFromIPs(ips)
			records, err := client.List(cmd.Context(), domain)
			if err != nil {
				return oops.Wrap(err)
			}
			recordSet := StringSetFromRecords(records)
			keeps := []dns.RecordResponse{}
			deletes := []dns.RecordResponse{}
			for _, record := range records {
				if ipSet.Contains(record.Content) {
					keeps = append(keeps, record)
				} else {
					deletes = append(deletes, record)
				}
			}
			posts := []netip.Addr{}
			for _, ip := range ips {
				if !recordSet.Contains(ip.String()) {
					posts = append(posts, ip)
				}
			}
			if (len(deletes) == 0) && (len(posts) == 0) {
				return nil
			}
			comment, err := config.Comment()
			if err != nil {
				return oops.Wrap(err)
			}
			resp, err := client.Update(cmd.Context(), comment, domain, deletes, posts)
			if err != nil {
				return oops.Wrap(err)
			}
			err = bot.Update(keeps, resp.Deletes, resp.Posts)
			if err != nil {
				return oops.Wrap(err)
			}
			return nil
		},
	}
	return cmd
}

func StringSetFromIPs(ips []netip.Addr) *hashset.Set {
	set := hashset.New()
	for _, ip := range ips {
		set.Add(ip.String())
	}
	return set
}

func StringSetFromRecords(records []dns.RecordResponse) *hashset.Set {
	set := hashset.New()
	for _, record := range records {
		set.Add(record.Content)
	}
	return set
}
