package utils

import (
    "net"
)

func ReadU32(b []byte) uint32 { // 4
    return uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24
}

func IPToU32(ipv4 string) uint32 {
    if ipBytes := net.ParseIP(ipv4).To4(); ipBytes != nil {
        return ReadU32(ipBytes)
    }
    return 0
}

func U32ToIP(ipv4 uint32) string {
    ip := net.IPv4(byte(ipv4), byte(ipv4>>8), byte(ipv4>>16), byte(ipv4>>24))
    return ip.String()
}