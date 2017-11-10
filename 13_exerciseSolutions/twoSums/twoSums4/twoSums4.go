package twoSums4

func TwoSum4(nums []int, target int) []int {

	if len(nums) < 2{
		return nil
	}

	m := make(map[int]int)

	for i, num := range nums {
		if j, exists := m[target-num]; exists {
			return []int{j,i}
		}
		m[num] = i
	}
	return nil
}

