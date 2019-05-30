import "strconv"

func str2tree(s string) *TreeNode {
	if 0 == len(s) {
		return nil
	}
	st := []*TreeNode{}
	for i := 0; i < len(s); i++ {
		if s[i] == ')' {
			st = st[:len(st)-1]
		} else if (s[i] >= '0' && s[i] <= '9') || s[i] == '-' {
			j := i
			for i+1 < len(s) && s[i+1] >= '0' && s[i+1] <= '9' {
				i++
			}
			v, _ := strconv.Atoi(s[j : i+1])
			n := &TreeNode{Val: v}
			if 0 != len(st) {
				t := st[len(st)-1]
				if nil == t.Left {
					t.Left = n
				} else {
					t.Right = n
				}
			}
			st = append(st, n)
		}
	}
	return st[len(st)-1]
}
