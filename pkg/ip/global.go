package ip

import "net/netip"

var reservedPrefixes = [...]netip.Prefix{
	// ref: <https://en.wikipedia.org/wiki/List_of_reserved_IP_addresses>
	netip.MustParsePrefix("100.64.0.0/10"), // Tailscale
	netip.MustParsePrefix("198.18.0.0/15"), // mihomo TUN
}

func IsGlobal(ip netip.Addr) bool {
	ip = ip.Unmap()
	if ip.IsPrivate() || !ip.IsGlobalUnicast() {
		return false
	}
	for _, prefix := range reservedPrefixes {
		if prefix.Contains(ip) {
			return false
		}
	}
	return true
}
