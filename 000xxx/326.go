var m map[int]bool = map[int]bool{
	1:                   true,
	3:                   true,
	9:                   true,
	27:                  true,
	81:                  true,
	243:                 true,
	729:                 true,
	2187:                true,
	6561:                true,
	19683:               true,
	59049:               true,
	177147:              true,
	531441:              true,
	1594323:             true,
	4782969:             true,
	14348907:            true,
	43046721:            true,
	129140163:           true,
	387420489:           true,
	1162261467:          true,
	3486784401:          true,
	10460353203:         true,
	31381059609:         true,
	94143178827:         true,
	282429536481:        true,
	847288609443:        true,
	2541865828329:       true,
	7625597484987:       true,
	22876792454961:      true,
	68630377364883:      true,
	205891132094649:     true,
	617673396283947:     true,
	1853020188851841:    true,
	5559060566555523:    true,
	16677181699666569:   true,
	50031545098999707:   true,
	150094635296999121:  true,
	450283905890997363:  true,
	1350851717672992089: true,
	4052555153018976267: true,
}

func isPowerOfThree(n int) bool {
	_, ok := m[n]
	return ok
}
