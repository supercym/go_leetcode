package main

import (
	"fmt"
)

func compress(chars []byte) int {
	start, end := 0, 0
	for {
		if chars[start] == chars[end] {
			end++
			if end != len(chars) {
				continue
			}
		}
		length := end - start
		if length > 1 {
			tmp := []byte{}
			for length != 0 {
				tmp = append(tmp, byte(length%10+48))
				length /= 10
			}
			for i, j := 0, len(tmp)-1; i < j; i, j = i+1, j-1 {
				tmp[i], tmp[j] = tmp[j], tmp[i]
			}
			// chars[start+1 : end] = tmp[:]
			tmpslice := chars[end:]
			chars = append(chars[:start+1], tmp...)
			chars = append(chars, tmpslice...)
			end = start + 1 + len(tmp)
		}
		start = end
		if end == len(chars) {
			break
		}
	}
	fmt.Printf("%c\n", chars)
	return len(chars)
}

func main() {
	chars1 := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	chars2 := []byte{'a'}
	chars3 := []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b'}
	fmt.Println(compress(chars1))
	fmt.Println(compress(chars2))
	fmt.Println(compress(chars3))
}
