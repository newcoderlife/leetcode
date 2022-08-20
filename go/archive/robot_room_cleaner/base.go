package robot_room_cleaner

type Robot struct{}

func (robot *Robot) Move() bool { return true }
func (robot *Robot) TurnLeft()  {}
func (robot *Robot) TurnRight() {}
func (robot *Robot) Clean()     {}

var (
	visited   map[int64]bool
	direction = [][]int64{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	r         *Robot
)

func key(x, y int64) int64 {
	return (x << 32) + y
}

func dfs(x, y int64, d int) {
	visited[key(x, y)] = true
	r.Clean()

	for i := 0; i < 4; i++ {
		di := (d + i) % 4
		nx, ny := x+direction[di][0], y+direction[di][1]
		if !visited[key(nx, ny)] && r.Move() {
			dfs(nx, ny, di)
		}

		r.TurnRight()
	}

	r.TurnLeft()
	r.TurnLeft()
	r.Move()
	r.TurnRight()
	r.TurnRight()
}

func cleanRoom(robot *Robot) {
	r = robot
	visited = make(map[int64]bool)
	dfs(0, 0, 0)
}
