package main

import "fmt"

func longestPalindrome(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}

	f := make([][]bool, n)
	for i := range f {
		f[i] = make([]bool, n)
		f[i][i] = true
	}

	start, maxLen := 0, 1

	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				if j-i == 1 || f[i+1][j-1] {
					f[i][j] = true
					if j-i+1 > maxLen {
						start = i
						maxLen = j - i + 1
					}
				}
			}
		}
	}
	return s[start : start+maxLen]
}

func main() {
	fmt.Println(longestPalindrome("babad")) // bab or aba
	fmt.Println(longestPalindrome("cbbd"))  // bb
}
