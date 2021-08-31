type AuthenticationManager struct {
	ttl int
	m   map[string]int
}

func Constructor(timeToLive int) AuthenticationManager {
	return AuthenticationManager{timeToLive, map[string]int{}}
}

func (this *AuthenticationManager) Generate(tokenId string, currentTime int) {
	this.m[tokenId] = currentTime + this.ttl
}

func (this *AuthenticationManager) Renew(tokenId string, currentTime int) {
	if this.m[tokenId] > currentTime {
		this.m[tokenId] = currentTime + this.ttl
	}
}

func (this *AuthenticationManager) CountUnexpiredTokens(currentTime int) int {
	o := 0
	for _, t := range this.m {
		if t > currentTime {
			o++
		}
	}
	return o
}
