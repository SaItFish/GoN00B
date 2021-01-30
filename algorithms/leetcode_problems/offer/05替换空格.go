package offer

import "strings"

func replaceSpace05(s string) string {
	return strings.Replace(s, " ", "%20", -1)
}
