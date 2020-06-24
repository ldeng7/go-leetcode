const MOD = 1000000007

func findGoodStrings(n int, s1 string, s2 string, evil string) int {
	t1 := [55][26]int{}
	t2 := [55][505]int{}
	le := len(evil)

	cal := func(s string) int {
		o, o1, l := 0, 0, len(s)
		for i := 0; i < l; i++ {
			o = (o*26 + int(s[i]-'a')) % MOD
		}
		for i, j := 0, 0; i < l; i++ {
			ke := int(s[i] - 'a')
			for k := 0; k < ke; k++ {
				t := t1[j][k]
				o1 = (o1 + t2[t][n-i-1+t]) % MOD
			}
			j = t1[j][ke]
		}
		return (o - o1 + MOD) % MOD
	}

	t := make([]int, le)
	for i := 1; i < le; i++ {
		j := i - 1
		for j >= 0 && evil[i] != evil[t[j]] {
			j = t[j] - 1
		}
		if j < 0 {
			t[i] = 0
		} else {
			t[i] = t[j] + 1
		}
	}

	e := evil[0]
	for i := byte(0); i < 26; i++ {
		if i+'a' == e {
			t1[0][i] = 1
		}
		t1[le][i] = le
	}
	for i := 1; i < le; i++ {
		e = evil[i]
		for j := byte(0); j < 26; j++ {
			if j+'a' == e {
				t1[i][j] = i + 1
			} else {
				t1[i][j] = t1[t[i-1]][j]
			}
		}
	}

	for i := 0; i <= le; i++ {
		t2[i][le] = 1
	}
	for i := le + 1; i <= n; i++ {
		t2[le][i] = t2[le][i-1] * 26 % MOD
		for j := le - 1; j >= 0; j-- {
			for k := 0; k < 26; k++ {
				p := t1[j][k]
				t2[j][i] = (t2[j][i] + t2[p][p+i-j-1]) % MOD
			}
		}
	}

	o := (cal(s2) - cal(s1) + 1 + MOD) % MOD
	for i := 0; i <= n-le; i++ {
		if s2[i:i+le] == evil {
			o = (o - 1 + MOD) % MOD
			break
		}
	}
	return o
}
