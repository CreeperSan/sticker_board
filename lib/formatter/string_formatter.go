package Formatter

import (
	"crypto/md5"
	"encoding/hex"
	"regexp"
	"strings"
)

func CheckStringWithLength(content string, minLength int, maxLength int) bool {
	contentLength := strings.Count(content, "")
	if contentLength < minLength || contentLength > maxLength {
		return false
	}
	return true
}

func CheckStringIsValidEmail(content string) bool {
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(content)
}

func FormatPassword(password string) string {
	h := md5.New()
	h.Write([]byte(password))
	return hex.EncodeToString(h.Sum(nil))
}
