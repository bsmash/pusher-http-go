package pusher

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"

	"strings"
)

func hmacSignature(to_sign, secret string) string {
	_auth_signature := hmac.New(sha256.New, []byte(secret))
	_auth_signature.Write([]byte(to_sign))
	return hex.EncodeToString(_auth_signature.Sum(nil))
}

func checkSignature(result, body, secret string) bool {
	expected := hmacSignature(body, secret)
	return result == expected
}

func createAuthMap(key, secret, string_to_sign string) map[string]string {
	auth_signature := hmacSignature(string_to_sign, secret)
	auth_string := strings.Join([]string{key, auth_signature}, ":")
	return map[string]string{"auth": auth_string}
}

func md5Signature(body []byte) string {
	_body_md5 := md5.New()
	_body_md5.Write([]byte(body))
	return hex.EncodeToString(_body_md5.Sum(nil))
}