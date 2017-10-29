package ipmath

import (
	"net"
	"testing"
)

func TestNextIP(t *testing.T) {
	ip := net.IP([]byte{1, 2, 3, 4})
	ip = NextIP(ip)
	if !ip.Equal(net.IP([]byte{1, 2, 3, 5})) {
		t.Fatal("next ip fail")
	}
}

func TestIsBroadcast1(t *testing.T) {
	ip, nw, _ := net.ParseCIDR("192.168.76.255/24")
	if !IsBroadcastAddress(ip, nw) {
		t.Fatal("should be true")
	}
}

func TestIsBroadcast2(t *testing.T) {
	ip, nw, _ := net.ParseCIDR("192.168.76.23/32")
	if IsBroadcastAddress(ip, nw) {
		t.Fatal("should be false")
	}
}
