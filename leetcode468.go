package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	queryIP := "01.01.01.01"
	fmt.Println(validIPAddress(queryIP))
}

func validIPAddress(queryIP string) string {
	if isIPv4(queryIP) {
		return "IPv4"
	}
	if isIPv6(queryIP) {
		return "IPv6"
	}
	return "Neither"
}

func isIPv4(queryIP string) bool {
	ipv4 := strings.Split(queryIP, ".")
	if len(ipv4) != 4 {
		return false
	}
	for _, s := range ipv4 {
		num, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return false
		}
		if num < 0 || num > 255 {
			return false
		}
		if len(s) > 1 && s[0] == '0' {
			return false
		}
	}
	return true
}

func isIPv6(queryIP string) bool {
	ipv6 := strings.Split(queryIP, ":")
	if len(ipv6) != 8 {
		return false
	}
	for _, s := range ipv6 {
		if len(s) < 1 || len(s) > 4 {
			return false
		}
		_, err := strconv.ParseInt(s, 16, 64)
		if err != nil {
			return false
		}
	}
	return true
}
