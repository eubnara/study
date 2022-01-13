package leetcode

func leftTrap(height []int, blackCount []int, rightIdx int) int {
	ret := 0
	if rightIdx <= 1 {
		return ret
	}
	highest := 0
	idx := -1
	for i := 0; i < rightIdx; i++ {
		v := height[i]
		if highest <= v {
			highest = v
			idx = i
		}
	}
	ret = highest*(rightIdx-idx-1) - (blackCount[rightIdx-1] - blackCount[idx])
	ret = ret + leftTrap(height, blackCount, idx)
	return ret
}

func rightTrap(height []int, blackCount []int, leftIdx int) int {
	ret := 0
	l := len(height)
	if leftIdx >= l-2 {
		return ret
	}
	highest := 0
	idx := -1
	for i := l - 1; i > leftIdx; i-- {
		v := height[i]
		if highest <= v {
			highest = v
			idx = i
		}
	}
	ret = highest*(idx-leftIdx-1) - (blackCount[idx-1] - blackCount[leftIdx])
	ret = ret + rightTrap(height, blackCount, idx)
	return ret
}

func trap(height []int) int {
	blackCount := make([]int, len(height))
	highest := 0
	highestIdx := -1
	ans := 0

	for i, v := range height {
		if highest < v {
			highest = v
			highestIdx = i
		}
		blackCount[i] = v
		if i > 0 {
			blackCount[i] += blackCount[i-1]
		}
	}
	if highestIdx == -1 {
		return ans
	}

	ans = leftTrap(height, blackCount, highestIdx) + rightTrap(height, blackCount, highestIdx)
	return ans
}
