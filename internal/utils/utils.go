package utils

import "strings"

func LeftNullTrimer(s string) string {
	return strings.TrimLeft(s, "\x00")
}
