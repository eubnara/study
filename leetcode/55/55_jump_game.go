package leetcode

func canJump(nums []int) bool {
	l := len(nums)
	maxIdx := 0
	for i := 0; i < l; i++ {
		v := nums[i]
		if maxIdx < i {
			return false
		}
		if maxIdx < i+v {
			maxIdx = i + v
		}
		if maxIdx >= l-1 {
			return true
		}
	}
	return false
}
