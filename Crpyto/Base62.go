package Crypto

import (
	"strings"
)

func decimalToBase62(n int) string {
	const charset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var result strings.Builder

	for n > 0 {
		remainder := n % 62
		result.WriteByte(charset[remainder])
		n /= 62
	}

	runes := []rune(result.String())
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
