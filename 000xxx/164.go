import "math"

func maximumGap(nums []int) int {
	if 0 == len(nums) {
		return 0
	}
	min, max := math.MaxInt64, math.MinInt64
	for _, n := range nums {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	sz := (max-min)/len(nums) + 1
	bucketMin, bucketMax := map[int]int{}, map[int]int{}

	for _, n := range nums {
		k := (n - min) / sz
		if v, ok := bucketMin[k]; !ok || n < v {
			bucketMin[k] = n
		}
		if v, ok := bucketMax[k]; !ok || n > v {
			bucketMax[k] = n
		}
	}

	o, pi := 0, 0
	for i := 1; i < len(nums); i++ {
		if min, ok := bucketMin[i]; ok {
			t := min - bucketMax[pi]
			if t > o {
				o = t
			}
			pi = i
		}
	}
	return o
}
