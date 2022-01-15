package main

import "fmt"

var cache [][]*bool

func isMatch(s string, p string) bool {
	cache = cache[:0]
	for i := 0; i < len(s); i++ {
		arr := make([]*bool, len(p))
		cache = append(cache, arr)
	}
	return isMatchRecursive(s, p, 0, 0)
}

func isMatchRecursive(s string, p string, sp, pp int) bool {
	fmt.Println(s[sp:], p[pp:])
	sl := len(s)
	pl := len(p)
	if sp < sl && pp < pl && cache[sp][pp] != nil {
		return *cache[sp][pp]
	}

	for sp < sl && pp < pl {
		if p[pp] == '*' {
			ans := isMatchRecursive(s, p, sp, pp+1) || isMatchRecursive(s, p, sp+1, pp)
			cache[sp][pp] = &ans
			return ans
		}
		if p[pp] == '?' || s[sp] == p[pp] {
			sp++
			pp++
			continue
		}
		return false
	}

	if sp >= sl {
		for pp < pl {
			if p[pp] != '*' {
				return false
			}
			pp++
		}
		return true
	}

	if sp < sl || pp < pl {
		return false
	}
	return true
}

func main() {
	// fmt.Println(isMatch("aa", "a"))                     // false
	// fmt.Println(isMatch("aa", "*"))                     // true
	// fmt.Println(isMatch("cb", "?a"))                    // false
	// fmt.Println(isMatch("acdcb", "a*c?b"))              // false
	// fmt.Println(isMatch("", "******"))                  // true
	// fmt.Println(isMatch("mississippi", "m??*ss*?i*pi")) // false
	// fmt.Println(isMatch("a", "aa"))                     // false
	// fmt.Println(isMatch("", "?"))                       // false
	// fmt.Println(isMatch("bcd", "??"))                   // false
	// fmt.Println(isMatch("ab", "?*"))                    // true
	// false
	fmt.Println(isMatch("abbabaaabbabbaababbabbbbbabbbabbbabaaaaababababbbabababaabbababaabbbbbbaaaabababbbaabbbbaabbbbababababbaabbaababaabbbababababbbbaaabbbbbabaaaabbababbbbaababaabbababbbbbababbbabaaaaaaaabbbbbaabaaababaaaabb", "**aa*****ba*a*bb**aa*ab****a*aaaaaa***a*aaaa**bbabb*b*b**aaaaaaaaa*a********ba*bbb***a*ba*bb*bb**a*b*bb"))

}
