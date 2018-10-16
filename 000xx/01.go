func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i, num := range nums {
		if j, ok := m[target-num]; ok {
			return []int{j, i}
		}
		m[num] = i
	}
	return nil
}
