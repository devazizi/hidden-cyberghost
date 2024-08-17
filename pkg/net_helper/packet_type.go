package net_helper

import "fmt"

func GetPacketType(packet []byte) {
	// Check if the packet is an IPv4 packet
	if packet[0]>>4 == 4 {
		// IPv4 Packet
		_ = int((packet[0] & 0x0F) * 4)
		protocol := packet[9] // Protocol field in IPv4 header

		fmt.Printf("IPv4 Packet Protocol: %d\n", protocol)

		switch protocol {
		case 1:
			fmt.Println("ICMP")
		case 6:
			fmt.Println("TCP")
		case 17:
			fmt.Println("UDP")
		default:
			fmt.Println("Other")
		}
	} else if packet[0]>>4 == 6 {
		// IPv6 Packet
		nextHeader := packet[6] // Next Header field in IPv6 header
		fmt.Printf("IPv6 Packet Next Header: %d\n", nextHeader)

		switch nextHeader {
		case 58:
			fmt.Println("ICMPv6")
		case 6:
			fmt.Println("TCP")
		case 17:
			fmt.Println("UDP")
		default:
			fmt.Println("Other")
		}
	} else {
		fmt.Println("Not an IPv4 or IPv6 packet")
	}
}
