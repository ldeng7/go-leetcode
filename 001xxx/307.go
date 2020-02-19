import "sort"

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func isSolvable(words []string, result string) bool {
	words = append(words, result)
	m := [26]int{}
	for i := 0; i < 26; i++ {
		m[i] = -1
	}
	n := 0
	for _, word := range words {
		for i := 0; i < len(word); i++ {
			if c := word[i] - 'A'; m[c] == -1 {
				m[c], n = n, n+1
			}
		}
	}

	cnts, v := make([]int, n), make([]bool, n)
	for i, word := range words {
		for j, k := len(word)-1, 1; j >= 0; j, k = j-1, k*10 {
			idx := m[word[j]-'A']
			if j == 0 {
				v[idx] = true
			}
			if i == len(words)-1 {
				cnts[idx] -= k
			} else {
				cnts[idx] += k
			}
		}
	}

	q := make([]int, 0, 32)
	for i := 0; i <= 9; i++ {
		q = append(q, i)
	}
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}
	sort.Slice(arr, func(i, j int) bool {
		return abs(cnts[arr[i]]) > abs(cnts[arr[j]])
	})
	arrm, sum := make([]int, n), 0
	for i := n - 1; i >= 0; i-- {
		sum += abs(cnts[arr[i]] * 9)
		arrm[i] = sum
	}

	var cal func(int, int, int) bool
	cal = func(i, sum, bits int) bool {
		if i == len(cnts) {
			return sum == 0
		} else if cnts[arr[i]] == 0 {
			return cal(i+1, sum, bits)
		} else if abs(sum) > arrm[i] {
			return false
		}
		for l, a := len(q), arr[i]; l > 0; l-- {
			t := q[0]
			q = q[1:]
			if (v[a] && t == 0) || bits&(1<<t) != 0 {
				q = append(q, t)
				continue
			}
			if cal(i+1, sum+cnts[a]*t, bits|(1<<t)) {
				return true
			}
			q = append(q, t)
		}
		return false
	}
	return cal(0, 0, 0)
}
