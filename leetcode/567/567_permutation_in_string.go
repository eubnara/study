package main

import "fmt"

func sameAlphaArr(a, b []int) bool {
	for i:=0;i<len(a);i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func checkInclusion(s1 string, s2 string) bool {
	b1, b2 := []byte(s1), []byte(s2)
	b1l, b2l := len(b1), len(b2)
	if b1l > b2l {
		return false
	}
	b1count := make([]int, 26)
	for _, v := range b1 {
		b1count[v-'a']++
	}
	b2count := make([]int, 26)
	for _, v := range b2[:b1l] {
		b2count[v-'a']++
	}
	if sameAlphaArr(b1count, b2count) {
		return true
	}

	for i:=b1l;i<b2l;i++ {
		prev, cur := b2[i-b1l]-'a', b2[i]-'a'
		if prev == cur {
			continue
		}
		b2count[prev]--
		b2count[cur]++
		if sameAlphaArr(b1count, b2count) {
			return true
		}
	}
	return false
}

func main() {
	// fmt.Println(checkInclusion("ab", "eidboaoo"))
	// fmt.Println(checkInclusion("adc", "dcda")) // true
	fmt.Println(checkInclusion("ab", "ba")) // true
}