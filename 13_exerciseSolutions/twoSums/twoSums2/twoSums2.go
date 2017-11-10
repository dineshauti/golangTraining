package twoSums2


func TwoSum2(nums []int, target int) []int {

	var arr []int

	var m map[int]int
	m = make(map[int]int)

	for i:=0; i<len(nums); i++ {
		m[nums[i]] = i
	}

	for j:=0; j<len(nums); j++ {
		complement := target - nums[j]

		if val, ok := m[complement]; ok {
			arr = append(arr, m[nums[j]], val)
			return arr
		}
	}
	return arr
}
