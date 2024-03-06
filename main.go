package main

import (
	"bytes"
	"fmt"
	"net"
	"os"
)

func leastCommonCIDR(ipAddresses []string) (string, error) {
	if len(ipAddresses) == 0 {
		return "", fmt.Errorf("no IP")
	}

	startIP := net.ParseIP(ipAddresses[0])
	endIP := net.ParseIP(ipAddresses[0])

	for _, ipStr := range ipAddresses[1:] {
		ip := net.ParseIP(ipStr)
		if bytes.Compare(ip, startIP) < 0 {
			startIP = ip
		} else if bytes.Compare(ip, endIP) > 0 {
			endIP = ip
		}
	}

	// find the mask from the start and end IP
	for mask := 32; mask >= 0; mask-- {
		maskedStart := net.CIDRMask(mask, 32)
		startNet := startIP.Mask(maskedStart)
		endNet := endIP.Mask(maskedStart)

		if net.IP.Equal(startNet, endNet) {
			return fmt.Sprintf("%s/%d", startNet.String(), mask), nil
		}
	}

	return "", fmt.Errorf("no common CIDR found")
}

func main() {
	ipAddresses := []string{
		"192.168.1.1",
		"192.168.1.2",
		"192.168.1.4",
		// "192.168.1.254",
	}

	lcc, err := leastCommonCIDR(ipAddresses)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	} else {
		fmt.Println("", lcc)
	}
}
