func entityParser(text string) string {
	bs := []byte(text)
	l := len(bs)
	bso, k, i := make([]byte, l), 0, 0
	for ; i < l-2; i++ {
		if b := bs[i]; b != '&' {
			bso[k], k = b, k+1
			continue
		}
		switch bs[i+1] {
		case 'a':
			switch bs[i+2] {
			case 'm':
				if i <= l-5 && bs[i+3] == 'p' && bs[i+4] == ';' {
					bso[k], k, i = '&', k+1, i+4
					continue
				}
			case 'p':
				if i <= l-6 && string(bs[i+3:i+6]) == "os;" {
					bso[k], k, i = '\'', k+1, i+5
					continue
				}
			}
		case 'f':
			if i <= l-7 && string(bs[i+2:i+7]) == "rasl;" {
				bso[k], k, i = '/', k+1, i+6
				continue
			}
		case 'g':
			if i <= l-4 && bs[i+2] == 't' && bs[i+3] == ';' {
				bso[k], k, i = '>', k+1, i+3
				continue
			}
		case 'l':
			if i <= l-4 && bs[i+2] == 't' && bs[i+3] == ';' {
				bso[k], k, i = '<', k+1, i+3
				continue
			}
		case 'q':
			if i <= l-5 && string(bs[i+2:i+6]) == "uot;" {
				bso[k], k, i = '"', k+1, i+5
				continue
			}
		}
		bso[k], k = '&', k+1
	}
	for ; i < l; i++ {
		bso[k], k = bs[i], k+1
	}
	return string(bso[:k])
}
