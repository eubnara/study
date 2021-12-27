package main

import "fmt"

func removeDupStr(s []string) []string {
	l := len(s)
	if l <= 1 {
		return s
	}
	out := s[:0]
	for i := 0; i < l-1; i++ {
		if s[i] == s[i+1] {
			continue
		}
		out = append(out, s[i])
	}
	out = append(out, s[l-1])
	return out
}

func removeDupStr2(s []string) []string {
	l := len(s)
	if l <= 1 {
		return s
	}
	out := s[:0]
	out = append(out, s[0])
	for i := 1; i < l; i++ {
		if s[i-1] == s[i] {
			continue
		}
		out = append(out, s[i])
	}
	return out
}

func main() {
	s := []string{"안", "안", "녕", "녕"}
	// s := []string{}
	fmt.Println(removeDupStr2(s))
	fmt.Println(s)
}
