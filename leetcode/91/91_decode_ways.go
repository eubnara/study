package leetcode


func numDecodings(s string) int {
	l := len(s)
	dp := make([]int, l)
	if s[0] == '0' {
		return 0
	} else {
		dp[0] = 1
	}
	if l == 1 {
		return dp[0]
	}
	if s[1] != '0' {
		dp[1] += 1
	}
	if s[0] - '0' == 1 {
		dp[1] += 1
	}
	if s[0] - '0' == 2 && s[1] - '0' <= 6 {
		dp[1] += 1
	}
	p, q := dp[0], dp[1]
	for i:=2;i<l;i++ {
		if p == 0 && q == 0 {
			return 0
		}
		t := s[i] - '0'
		t2 := s[i-1] - '0'
		if t != 0 {
			dp[i] += q
		}
		if t2 == 1  {
			dp[i] += p
		}
		if t2 == 2 && t <= 6{
			dp[i] += p
		}
		p, q = q, dp[i]
	}
	return dp[l-1]
}
