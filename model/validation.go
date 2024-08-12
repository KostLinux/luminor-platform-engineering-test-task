package model

import (
	"fmt"
	enum "messager-web-app/model/enum"
	"net"
	"strconv"
	"strings"
)

func ValidatePayload(payload Payload) error {
	if !isValidUnixTimestamp(payload.Ts) {
		return fmt.Errorf("invalid Unix timestamp")
	}
	if !enum.ValidSenders[payload.Sender] {
		return fmt.Errorf("invalid sender")
	}
	if payload.Message == (Message{}) {
		return fmt.Errorf("message cannot be empty")
	}
	if payload.SentFromIP != "" && !isValidIPv4(payload.SentFromIP) {
		return fmt.Errorf("invalid IPv4 address")
	}
	return nil
}

func isValidUnixTimestamp(ts string) bool {
	// Check if the timestamp is a valid Unix timestamp
	timestamp, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		return false
	}
	// Unix timestamp should be within a reasonable range
	// For example, between 1970-01-01 and 2038-01-19
	return timestamp >= 0 && timestamp <= 2147483647
}

func isValidIPv4(ip string) bool {
	// Check if the IP address is a valid IPv4 address
	validIP := net.ParseIP(ip)
	if validIP == nil {
		return false
	}

	// Ensure the IP address is in the correct format (x.y.z.f)
	parts := strings.Split(ip, ".")
	if len(parts) != 4 {
		return false
	}

	// Ensure each part of the IP address is a valid number between 0 and 255
	for _, part := range parts {
		if num, err := strconv.Atoi(part); err != nil || num < 0 || num > 255 {
			return false
		}
	}

	return true
}
