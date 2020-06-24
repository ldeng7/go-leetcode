var vowels = [...]byte{'a', 'e', 'i', 'o', 'u'}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func maxVowels(s string, k int) int {
	o, c, ar := 0, 0, [26]int{}
	for _, v := range vowels {
		ar[v-'a'] = 1
	}
	for i := 0; i < k; i++ {
		c += ar[s[i]-'a']
	}
	o = max(o, c)
	for i := k; i < len(s); i++ {
		c = c - ar[s[i-k]-'a'] + ar[s[i]-'a']
		o = max(o, c)
	}
	return o
}
