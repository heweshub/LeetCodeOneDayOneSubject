package main

func asteroidCollision(a []int) (st []int) {
	for _, as := range a {
		alive := true
		for alive && as < 0 && len(st) > 0 && st[len(st)-1] > 0 {
			alive = st[len(st)-1] < -as
			if st[len(st)-1] <= -as {
				st = st[:len(st)-1]
			}
		}
		if alive {
			st = append(st, as)
		}
	}
	return
}
