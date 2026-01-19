package main

func averageCalc(nums ...int) float32 {
	sum := 0

	if len(nums) == 0 {
		return 0
	}

	for _, num := range nums {
		sum += num
	}

	// Calculate average as float32
	return float32(sum) / float32(len(nums))
}