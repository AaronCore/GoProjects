package split_string

import "strings"

// SplitString 字符串切割
func SplitString(s, sep string) []string {
	// var ret []string
	ret := make([]string, 0, strings.Count(s, sep)+1)
	index := strings.Index(s, sep)
	for index > 0 {
		ret = append(ret, s[:index])
		s = s[index+1:]
		index = strings.Index(s, sep)
	}
	ret = append(ret, s)
	return ret
}
