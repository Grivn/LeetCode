package lc14

import (
	"fmt"
	"testing"
)

func longestCommonPrefix(strs []string) string {
	minLen := len(strs[0])

	for _, str := range strs {
		if len(str) < minLen {
			minLen = len(str)
		}
	}

	prefix := ""
	for i:=0; i<minLen; i++ {
		flag := true
		curPrefix := strs[0][:i+1]
		for _, str := range strs {
			if str[:i+1] != curPrefix {
				flag = false
				break
			}
		}
		if flag {
			prefix = curPrefix
		} else {
			break
		}
	}

	return prefix
}

func TestLC14(t *testing.T) {
	strs := []string{"flower","flow","flight"}
	fmt.Println(longestCommonPrefix(strs))
}
