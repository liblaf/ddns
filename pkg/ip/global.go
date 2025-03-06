package ip

import "net/netip"

func IsGlobal(ip netip.Addr) bool {
	return ip.IsValid() &&
		!(ip.IsInterfaceLocalMulticast() ||
			ip.IsLinkLocalMulticast() ||
			ip.IsLinkLocalUnicast() ||
			ip.IsLoopback() ||
			ip.IsPrivate())
}
