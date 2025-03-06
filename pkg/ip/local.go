package ip

import (
	"net"
	"net/netip"

	"github.com/rs/zerolog/log"
	"github.com/samber/oops"
)

func GlobalIPs() ([]netip.Addr, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, oops.Wrap(err)
	}
	results := make([]netip.Addr, 0, len(addrs))
	for _, addr := range addrs {
		ipNet := addr.(*net.IPNet)
		ip, ok := AddrFromIPNet(*ipNet)
		if !ok {
			log.Warn().Any("ipNet", ipNet).Send()
			continue
		}
		if IsGlobal(ip) {
			results = append(results, ip)
		}
	}
	return results, nil
}

func AddrFromIPNet(ipNet net.IPNet) (ip netip.Addr, ok bool) {
	if ipv4 := ipNet.IP.To4(); ipv4 != nil {
		var ok bool
		ip, ok = netip.AddrFromSlice(ipv4)
		if !ok {
			return netip.Addr{}, false
		}
	} else if ipv6 := ipNet.IP.To16(); ipv6 != nil {
		var ok bool
		ip, ok = netip.AddrFromSlice(ipv6)
		if !ok {
			return netip.Addr{}, false
		}
	} else {
		return netip.Addr{}, false
	}
	return ip, true
}
