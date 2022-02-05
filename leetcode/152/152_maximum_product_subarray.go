package main

import "fmt"

func maxProduct(nums []int) int {
	ret := nums[0]

	pMax := 0
	nMax := 0
	for i := 0; i < len(nums); i++ {
		v := nums[i]
		updated := false
		if v == 0 {
			pMax = 0
			nMax = 0
			updated = true
		}
		if v > 0 {
			updated = true
			if pMax == 0 {
				pMax = v
			} else {
				pMax *= v
			}

			if nMax != 0 {
				nMax *= v
			}
		}
		if v < 0 {
			tmpPMax, tmpNMax := pMax, nMax
			if pMax == 0 {
				tmpNMax = v
			} else {
				tmpNMax = pMax * v
			}
			if nMax == 0 {
				tmpPMax = 0
			} else {
				tmpPMax = nMax * v
				updated = true
			}
			pMax, nMax = tmpPMax, tmpNMax
		}
		if updated && pMax > ret {
			ret = pMax
		}
	}

	return ret
}

func main() {
	// fmt.Println(maxProduct([]int{-2})) // -2
	// fmt.Println(maxProduct([]int{-3, -1, -1})) // 3
	fmt.Println(maxProduct([]int{0, 2})) // 2
}
