func findInMountainArray(target int, mountainArr *MountainArray) int {
	n := mountainArr.length()
	l, r, hi := 0, n-1, -1
	for l < r {
		m := l + (r-l)>>1
		if m == 0 {
			m += 1
			r += 1
		}
		if m == n-1 {
			l -= 1
			m -= 1
		}

		vl, vm, vr := mountainArr.get(m-1), mountainArr.get(m), mountainArr.get(m+1)
		if vm > vl && m < vr {
			l = m + 1
			if vl == target {
				return m - 1
			}
			if vm == target {
				return m
			}
			if vr == target {
				return m + 1
			}
		} else if vm < vl && vm > vr {
			r = m - 1
		} else {
			hi = m
			if vm == target {
				return m
			}
			break
		}
	}
	if hi == -1 {
		hi = l
	}

	l, r = 0, hi
	for l <= r {
		m := l + (r-l)>>1
		vm := mountainArr.get(m)
		if target == vm {
			return m
		} else if target > vm {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	l, r = hi, n-1
	for l <= r {
		m := l + (r-l)>>1
		vm := mountainArr.get(m)
		if target == vm {
			return m
		} else if target < vm {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return -1
}
