package config

import (
	"os"
	"strings"

	"github.com/samber/oops"
	"github.com/spf13/viper"
)

func Comment() (string, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return "", oops.Wrap(err)
	}
	return hostname, nil
}

func Domain() (string, error) {
	domain := viper.GetString("domain")
	if domain != "" {
		return domain, nil
	}
	hostname, err := os.Hostname()
	if err != nil {
		return "", oops.Wrap(err)
	}
	domain = hostname + ".ddns.liblaf.me"
	domain = strings.ToLower(domain)
	return domain, nil
}
