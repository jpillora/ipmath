package ipmath

import (
	"crypto/sha1"
	"encoding/binary"
	"math"
	"net"
)

//ToUInt32 converts an IPv4 address into
//a uint32
func ToUInt32(ip net.IP) uint32 {
	buff := make([]byte, 4)
	copy(buff, []byte(ip))
	return binary.BigEndian.Uint32(buff)
}

//FromUInt32 converts a uint32 into
//an IPv4 address
func FromUInt32(u uint32) net.IP {
	buff := make([]byte, 4)
	binary.BigEndian.PutUint32(buff, u)
	return net.IP(buff)
}

//DeltaIP returns the IPv4 delta-many places away
func DeltaIP(ip net.IP, delta int) net.IP {
	if delta == 0 {
		return ip
	}
	i := int64(ToUInt32(ip))
	i += int64(delta)
	if i > math.MaxUint32 {
		i = math.MaxUint32
	}
	return FromUInt32(uint32(i))

}

//NextIP returns the next IPv4 in sequence
func NextIP(ip net.IP) net.IP {
	return DeltaIP(ip, 1)
}

//PrevIP returns the previous IPv4 in sequence
func PrevIP(ip net.IP) net.IP {
	return DeltaIP(ip, 1)
}

//IsNetworkAddress returns whether the given IPv4 address
//is the network address of the given IPv4 subnet
func IsNetworkAddress(ip net.IP, network *net.IPNet) bool {
	curr := binary.BigEndian.Uint32([]byte(ip))
	mask := binary.BigEndian.Uint32([]byte(network.Mask))
	if mask == math.MaxUint32 {
		return false
	}
	return (^mask & curr) == uint32(0)
}

//IsBroadcastAddress returns whether the given IPv4 address
//is the broadcast address of the given IPv4 subnet
func IsBroadcastAddress(ip net.IP, network *net.IPNet) bool {
	curr := binary.BigEndian.Uint32([]byte(ip))
	mask := binary.BigEndian.Uint32([]byte(network.Mask))
	if mask == math.MaxUint32 {
		return false
	}
	return (mask | curr) == math.MaxUint32
}

//NetworkSize returns the number of addresses in a subnet
func NetworkSize(network *net.IPNet) uint32 {
	mask := binary.BigEndian.Uint32([]byte(network.Mask))
	return ^mask
}

//Hash an IP with SHA1
func Hash(ip net.IP) []byte {
	input := []byte(ip.To4())
	h := sha1.New()
	h.Write(input)
	output := h.Sum(nil)
	return output
}
