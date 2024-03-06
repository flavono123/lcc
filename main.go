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

	startIP, err := parseIPString(ipAddresses[0])
	if err != nil {
		return "", err
	}
	endIP, err := parseIPString(ipAddresses[0])
	if err != nil {
		return "", err
	}

	for _, ipStr := range ipAddresses[1:] {
		ip, err := parseIPString(ipStr)
		if err != nil {
			return "", err
		}
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

func parseIPString(ipStr string) (net.IP, error) {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return nil, fmt.Errorf("invalid IP address: %s", ipStr)
	}
	return ip, nil
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
