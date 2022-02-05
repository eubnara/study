package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
    found := map[int]int{}
    ans := 0
    sum := 0
    for i, c := range s {
        n := int(c - 'a')
        if idx, exists := found[n];exists {
            if sum > ans {
                ans = sum
            }
			tmp := i - idx
			if sum > tmp {
				sum = tmp
			} else {
				sum++
			}
            found[n] = i
			continue
        }
		found[n] = i
        sum++
    }
    if sum > ans {
        ans = sum
    }
    return ans
}

func main() {
	// fmt.Println(lengthOfLongestSubstring("abcabcbb")) // 3
	// fmt.Println(lengthOfLongestSubstring("abba")) // 2
	// fmt.Println(lengthOfLongestSubstring("tmmzuxt")) // 5
	fmt.Println(lengthOfLongestSubstring("abcabcbb")) // 3
}