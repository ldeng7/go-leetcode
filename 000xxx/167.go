import "sort"

func twoSum(numbers []int, target int) []int {
	l := len(numbers)
	for i, num := range numbers {
		j := sort.Search(l, func(k int) bool {
			return numbers[k] >= target-num
		})
		if j != i && j < l && numbers[j] == target-num {
			if i < j {
				return []int{i + 1, j + 1}
			}
			return []int{j + 1, i + 1}
		}
	}
	return nil
}
