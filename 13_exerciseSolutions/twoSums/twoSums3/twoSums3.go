package twoSums3

func TwoSum3(nums []int, target int) []int {

	var arr []int

	//var m map[int]int
	m := make(map[int]int)

	// If there are only two elements do not apply the mop logic, since it increases the overhead
	if len(nums) == 2 {
		if nums[0] + nums[1] == target {
			arr = append(arr,0,1)
			return arr
		}
	}

	for j:=0; j<len(nums); j++ {
		complement := target - nums[j]

		if val, ok := m[complement]; ok {
			arr = append(arr, val, j)
			return arr
		}

		m[nums[j]] = j
	}

	return arr
}
