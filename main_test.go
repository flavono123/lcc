package main

import (
	"fmt"
	"net"
	"reflect"
	"testing"
)

func TestLeastCommonCIDR(t *testing.T) {
	testCases := []struct {
		ipAddresses  []string
		expectedCIDR string
		expectedErr  error
	}{
		// Test case 1: Single IP address
		{
			ipAddresses:  []string{"192.168.0.1"},
			expectedCIDR: "192.168.0.1/32",
			expectedErr:  nil,
		},
		// Test case 2: Multiple IP addresses with common CIDR
		{
			ipAddresses:  []string{"192.168.0.1", "192.168.0.2", "192.168.0.3"},
			expectedCIDR: "192.168.0.0/30",
			expectedErr:  nil,
		},
		// Test case 3: Empty IP addresses
		{
			ipAddresses:  []string{},
			expectedCIDR: "",
			expectedErr:  fmt.Errorf("no IP"),
		},
		// Test case 4: Invalid IP address
		{
			ipAddresses:  []string{"192.168.0.1", "invalid-ip"},
			expectedCIDR: "",
			expectedErr:  fmt.Errorf("invalid IP address: invalid-ip"),
		},
	}

	for _, tc := range testCases {
		cidr, err := leastCommonCIDR(tc.ipAddresses)

		if cidr != tc.expectedCIDR {
			t.Errorf("Expected CIDR: %s, but got: %s", tc.expectedCIDR, cidr)
		}

		if fmt.Sprint(err) != fmt.Sprint(tc.expectedErr) {
			t.Errorf("Expected error: %v, but got: %v", tc.expectedErr, err)
		}
	}
}
func TestParseIPString(t *testing.T) {
	testCases := []struct {
		ipStr       string
		expectedIP  net.IP
		expectedErr error
	}{
		// Test case 1: Valid IP address
		{
			ipStr:       "192.168.0.1",
			expectedIP:  net.ParseIP("192.168.0.1"),
			expectedErr: nil,
		},
		// Test case 2: Invalid IP address
		{
			ipStr:       "invalid-ip",
			expectedIP:  nil,
			expectedErr: fmt.Errorf("invalid IP address: invalid-ip"),
		},
	}

	for _, tc := range testCases {
		ip, err := parseIPString(tc.ipStr)

		if !reflect.DeepEqual(ip, tc.expectedIP) {
			t.Errorf("Expected IP: %v, but got: %v", tc.expectedIP, ip)
		}

		if fmt.Sprint(err) != fmt.Sprint(tc.expectedErr) {
			t.Errorf("Expected error: %v, but got: %v", tc.expectedErr, err)
		}
	}
}
