func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	color := image[sr][sc]
	if color == newColor {
		return image
	}
	var cal func(int, int)
	cal = func(i, j int) {
		if i < 0 || i >= len(image) || j < 0 || j >= len(image[0]) || image[i][j] != color {
			return
		}
		image[i][j] = newColor
		cal(i+1, j)
		cal(i, j+1)
		cal(i-1, j)
		cal(i, j-1)
	}
	cal(sr, sc)
	return image
}
