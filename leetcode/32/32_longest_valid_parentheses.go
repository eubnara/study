package leetcode

/*

- longest 이기 때문에 연결되어 있어야 한다.
- 여러짝이 생길 수는 없다.


- 첫 () 를 찾는다.
- 양끝을 인덱스로 퍼져나가며 찾는다. 오른쪽에서 ( 가 나오면

*/

func longestValidParentheses(s string) int {
	found := make([]bool, len(s))
	ret := 0
	l := len(s)

	// 첫 () 찾기
	for i := 0; i < l-1; i++ {
		if s[i] == '(' && s[i+1] == ')' {
			found[i] = true
			found[i+1] = true
			i++
			continue
		}
	}

	for i := 0; i < l-1; i++ {
		if !found[i] {
			continue
		}
		left := i
		right := i + 1
		for {
			if right+1 < l && found[right+1] {
				right++
				continue
			}
			if left-1 >= 0 && found[left-1] {
				left--
				continue
			}
			if left-1 < 0 || right+1 >= l {
				break
			}
			if s[left-1] == '(' && s[right+1] == ')' {
				found[left-1] = true
				found[right+1] = true
				left--
				right++
				continue
			} else {
				break
			}
		}
		tmp := right - left + 1
		if ret < tmp {
			ret = tmp
		}
		i = right + 1
	}
	return ret
}
