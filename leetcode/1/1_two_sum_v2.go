package leetcode

func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i, v := range nums {
		complement := target - v
		if j, exists := m[complement]; exists {
			return []int{i, j}
		}
		m[v] = i
	}
	return []int{}
}
