package cmath

func Max[T RealNumber](a, b T) T {
	if a < b {
		return b
	}
	return a
}

func MaxInSlice[T RealNumber](nums ...T) T {
	var res T
	if len(nums) == 0 {
		return res
	}
	res = nums[0]
	for _, num := range nums {
		if res < num {
			res = num
		}
	}
	return res
}

func Min[T RealNumber](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func MinInSlice[T RealNumber](nums ...T) T {
	var res T
	if len(nums) == 0 {
		return res
	}
	res = nums[0]
	for _, num := range nums {
		if res > num {
			res = num
		}
	}
	return res
}
