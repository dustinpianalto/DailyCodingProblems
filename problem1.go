package main

func problem1(val int, nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	comp := make(map[int]struct{})

	for _, num := range nums {
		if _, ok := comp[val-num]; ok {
			return true
		}
		comp[num] = struct{}{}
	}

	return false
}
