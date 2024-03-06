package main

import (
	"bytes"
	"fmt"
	"net"
)

func leastCommonCIDR(ipAddresses []string) (string, error) {
	if len(ipAddresses) == 0 {
		return "", fmt.Errorf("IP 주소 목록이 비어 있습니다")
	}

	// 첫 번째 IP 주소로 시작합니다.
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

	// 시작 IP와 종료 IP 사이의 네트워크 마스크를 찾습니다.
	for mask := 32; mask >= 0; mask-- {
		maskedStart := net.CIDRMask(mask, 32)
		startNet := startIP.Mask(maskedStart)
		endNet := endIP.Mask(maskedStart)

		if net.IP.Equal(startNet, endNet) {
			return fmt.Sprintf("%s/%d", startNet.String(), mask), nil
		}
	}

	return "", fmt.Errorf("적절한 CIDR을 계산할 수 없습니다")
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
		fmt.Println("에러:", err)
	} else {
		fmt.Println("최소 공통 CIDR:", lcc)
	}
}
