package base64

import (
	"encoding/base64"
	"errors"
	"strings"
)

// Base64Encoding []byte序列化成base64string
func Base64Encoding(b []byte) string {
	if len(b) == 0 {
		return ""
	}
	return base64.StdEncoding.EncodeToString(b)
}

// Base64UrlEncodeing []byte序列化成base64string
func Base64UrlEncodeing(b []byte) string {
	if len(b) == 0 {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

// Base64Decoding string 反序列化成[]byte
func Base64Decoding(s string) ([]byte, error) {
	if len(s) == 0 {
		return nil, errors.New("src must be not null")
	}
	if l := len(s) % 4; l != 0 {
		s += strings.Repeat("=", 4-l)
	}
	return base64.StdEncoding.DecodeString(s)
}

// Base64UrlDecoding string 反序列化成[]byte
func Base64UrlDecoding(s string) ([]byte, error) {
	if len(s) == 0 {
		return nil, errors.New("src must be not null")
	}
	if l := len(s) % 4; l != 0 {
		s += strings.Repeat("=", 4-l)
	}
	return base64.URLEncoding.DecodeString(s)
}
