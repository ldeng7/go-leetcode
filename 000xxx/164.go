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
	mMin, mMax := map[int]int{}, map[int]int{}

	for _, n := range nums {
		k := (n - min) / sz
		if v, ok := mMin[k]; !ok || n < v {
			mMin[k] = n
		}
		if v, ok := mMax[k]; !ok || n > v {
			mMax[k] = n
		}
	}

	o, pi := 0, 0
	for i := 1; i < len(nums); i++ {
		if min, ok := mMin[i]; ok {
			t := min - mMax[pi]
			if t > o {
				o = t
			}
			pi = i
		}
	}
	return o
}
