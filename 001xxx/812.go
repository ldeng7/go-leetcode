func squareIsWhite(coordinates string) bool {
	return (coordinates[0]-'a')&1 != (coordinates[1]-'1')&1
}
