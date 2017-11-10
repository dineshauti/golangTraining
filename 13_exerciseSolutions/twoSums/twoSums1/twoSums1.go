package twoSums1

func TwoSum1(nums []int, target int) []int {

	arr := []int{0, 0}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				if i != j {
					arr[0] = i
					arr[1] = j
					return arr
				}
			}
		}
	}
	return arr
}
