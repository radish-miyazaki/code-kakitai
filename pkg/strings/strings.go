package strings

import "strings"

func RemoveHyphen(s string) string {
	return strings.ReplaceAll(s, "-", "")
}
