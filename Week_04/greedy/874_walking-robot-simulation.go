package main

func robotSim(commands []int, obstacles [][]int) int {
	dx := []int{0, 1, 0, -1}
	dy := []int{1, 0, -1, 0}
	x, y, di, ans := 0, 0, 0, 0

	ma := make(map[int64]struct{}, len(obstacles))
	for _, obs := range obstacles {
		ox := int64(obs[0] + 30000)
		oy := int64(obs[1] + 30000)
		ma[(ox<<16)+oy] = struct{}{}
	}

	for _, cmd := range commands {
		if cmd == -2 {
			// 左转90度等于右转270度
			di = (di + 3) % 4
		} else if cmd == -1 {
			di = (di + 1) % 4
		} else {
			for k := 0; k < cmd; k++ {
				// 每迭代1次，往该方向前进1个单位
				nx, ny := x+dx[di], y+dy[di]
				ox, oy := int64(nx+30000), int64(ny+30000)
				code := (ox << 16) + oy
				if _, ok := ma[code]; !ok {
					x, y = nx, ny
					v := x*x + y*y
					if v > ans {
						ans = v
					}
				}
			}
		}
	}

	return ans
}