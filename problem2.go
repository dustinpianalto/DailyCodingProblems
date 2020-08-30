package main

func problem2(nums []int) []int {
	var out []int
	for i := range nums {
		val := multiplySlice(nums[:i]) * multiplySlice(nums[i+1:])
		out = append(out, val)
	}
	return out
}

func multiplySlice(nums []int) int {
	out := 1
	for _, num := range nums {
		out *= num
	}
	return out
}
