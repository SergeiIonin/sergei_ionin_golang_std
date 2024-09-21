package main

import "fmt"

func main() {
	in := "the sky is blue"
	res := reverseWords(in)
	fmt.Printf("Reversed string is %s", res)
}

// https://leetcode.com/problems/reverse-words-in-a-string/
func reverseWords(s string) string {
    runes := []rune(s)

	ind := 0
	for i, elem := range runes {
		if elem != ' ' {
			ind = i
			break
		}
	}
	runes = runes[ind:]

	r := len(runes)-1
	l := r-1
	acc := make([]rune, 0, 5)

	for r >= 0 {
		if runes[r] == ' ' {
			r--
			l = r-1
		}
		if l < 0 {
			if runes[r] != ' ' {
				acc = append(acc, runes[:r+1]...)
			}
			break
		}
		if runes[l] != ' ' {
			l--
		} else {
			if runes[r] != ' ' {
				acc = append(append(acc, runes[l+1:r+1]...), ' ')
			}
			r = l-1
			l = r-1
		}
	}

	return string(acc)
}
