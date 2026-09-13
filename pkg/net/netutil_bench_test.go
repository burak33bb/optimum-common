package net_test

import (
	stdnet "net"
	"strings"
	"testing"

	netpkg "github.com/getoptimum/optimum-common/pkg/net"
)

var (
	benchBool       bool
	benchString     string
	benchMultiAddrs any
)

func BenchmarkIPClassifiers(b *testing.B) {
	for _, tc := range []struct {
		name string
		ip   stdnet.IP
		fn   func(stdnet.IP) bool
	}{
		{name: "PrivateIPv4", ip: stdnet.ParseIP("10.1.2.3"), fn: netpkg.IsPrivateOrULA},
		{name: "PublicIPv4", ip: stdnet.ParseIP("8.8.8.8"), fn: netpkg.IsGlobalUnicast},
		{name: "ULAIPv6", ip: stdnet.ParseIP("fd00::1"), fn: netpkg.IsPrivateOrULA},
		{name: "PublicIPv6", ip: stdnet.ParseIP("2606:4700:4700::1111"), fn: netpkg.IsGlobalUnicast},
	} {
		b.Run(tc.name, func(b *testing.B) {
			b.ReportAllocs()
			for range b.N {
				benchBool = tc.fn(tc.ip)
			}
		})
	}
}

func BenchmarkGetIPProtocol(b *testing.B) {
	for _, ip := range []string{
		"192.0.2.10",
		"2001:db8::1",
		"not-an-ip",
	} {
		b.Run(ip, func(b *testing.B) {
			b.ReportAllocs()
			for range b.N {
				benchString = netpkg.GetIPProtocol(ip)
			}
		})
	}
}

func BenchmarkParseCloudflareTrace(b *testing.B) {
	body := "fl=abc\nh=bootstrap.getoptimum.io\nip=203.0.113.42\nts=1234567890\nvisit_scheme=https\n"
	b.SetBytes(int64(len(body)))
	b.ReportAllocs()
	b.ResetTimer()

	for range b.N {
		var err error
		benchString, err = netpkg.ParseCloudflareTrace(strings.NewReader(body))
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMultiAddressBuilder(b *testing.B) {
	for _, tc := range []struct {
		name string
		ip   stdnet.IP
	}{
		{name: "IPv4", ip: stdnet.ParseIP("192.0.2.10")},
		{name: "IPv6", ip: stdnet.ParseIP("2001:db8::1")},
	} {
		b.Run(tc.name, func(b *testing.B) {
			b.ReportAllocs()
			for range b.N {
				addrs, err := netpkg.MultiAddressBuilder(tc.ip, 30303)
				if err != nil {
					b.Fatal(err)
				}
				benchMultiAddrs = addrs
			}
		})
	}
}
