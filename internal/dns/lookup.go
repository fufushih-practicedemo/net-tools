package dns

import (
	"fmt"
	"net"
)

func Lookup(domain string) ([]net.IP, error) {
	ips, err := net.LookupIP(domain)

	if err != nil {
		return nil, fmt.Errorf("error looking up IP for %s: %v", domain, err)
	}

	return ips, nil
}
