func cleanRoom(robot *Robot) {
	dirs := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	m := map[int]map[int]bool{}
	var cal func(int, int, int)
	cal = func(y, x, dir int) {
		robot.Clean()
		if my, ok := m[y]; ok {
			my[x] = true
		} else {
			m[y] = map[int]bool{x: true}
		}
		for i := 0; i < 4; i++ {
			j := dir + i
			if j >= 4 {
				j -= 4
			}
			y1, x1 := y+dirs[j][0], x+dirs[j][1]
			if my, ok := m[y1]; (!ok || !my[x1]) && robot.Move() {
				cal(y1, x1, j)
				robot.TurnRight()
				robot.TurnRight()
				robot.Move()
				robot.TurnLeft()
				robot.TurnLeft()
			}
			robot.TurnRight()
		}
	}
	cal(0, 0, 0)
}
