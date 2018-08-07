package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	pre := ""

	for i := 0; i < len(strs)-1; i++ {
		if i == 0 {
			pre = CommonPre(strs[i], strs[i+1])
		} else {
			tmp := CommonPre(strs[i], strs[i+1])
			if len(pre) <= len(tmp) {
				continue
			} else {
				if tmp[:] == pre[:len(tmp)] {
					pre = tmp[:]
				} else {
					break
				}
			}

		}
	}
	return pre

}

func CommonPre(str1, str2 string) string {
	pre := []byte{}
	min_length := 0
	if len(str1) <= len(str2) {
		min_length = len(str1)
	} else {
		min_length = len(str2)
	}
	for i := 0; i < min_length; i++ {
		if str1[i] == str2[i] {
			pre = append(pre, str1[i])
		} else {
			break
		}
	}
	return string(pre)
}

func main() {
	s1 := []string{"flower", "flow", "flight"}
	s2 := []string{"dog", "racecar", "car"}
	s3 := []string{"hello", "hell", "heat"}
	fmt.Println(longestCommonPrefix(s1))
	fmt.Println(longestCommonPrefix(s2))
	fmt.Println(longestCommonPrefix(s3))

}
