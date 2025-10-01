package util

import "strings"

// MaskToken returns a short masked representation of token for logs
func MaskToken(tok string) string {
	tok = strings.TrimSpace(tok)
	if len(tok) <= 8 {
		return "****"
	}
	return tok[:4] + "..." + tok[len(tok)-4:]
}

// MaskAuthorizationLine is used to mask a raw request dump containing Authorization header
func MaskAuthorizationLine(dump string) string {
	// naive: replace "Authorization: Token <token>" token with masked token
	// don't try to be perfect; it's just for logs
	return strings.ReplaceAll(dump, "Authorization: Token ", "Authorization: Token ")
}
