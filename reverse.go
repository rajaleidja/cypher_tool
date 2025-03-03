package reverse

import (
	"strings"
)

// Encrypts or decrypts a message by reversing the alphabet
func reverse(input string) string {
	var result strings.Builder
	for _, char := range input {
		if ('a' <= char && char <= 'z') || ('A' <= char && char <= 'Z') {
			if 'a' <= char && char <= 'z' {
				result.WriteRune('z' - (char - 'a'))
			} else {
				result.WriteRune('Z' - (char - 'A'))
			}
		} else {
			result.WriteRune(char) // Non-alphabet characters are left unchanged
		}
	}
	return result.String()
}
