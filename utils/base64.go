package utils

import "encoding/base64"

// The base64Decode function decodes a base64 encoded string and returns the decoded string.
func Base64Decode(encoded string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return "", err
	}
	return string(decoded), nil
}

// The base64Encode function encodes a string to base64 and returns the encoded string.
func Base64Encode(decoded string) string {
	return base64.StdEncoding.EncodeToString([]byte(decoded))
}
