package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	prefix := strs[0]
	for _, v := range strs {
		for strings.Index(v,prefix) != 0 {
			if len(prefix) == 0 {
				return ""
			}
			prefix = prefix[:len(prefix)-1]
		}
	}
	return prefix
}

func main() {
	str := []string{"flower","flow","flight"}
	prefix := longestCommonPrefix(str)
	fmt.Println(prefix)
}
