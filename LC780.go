package main

func reachingPoints(sx int, sy int, tx int, ty int) bool {
	// 反向计算，这样会更快，提示sx和sy都大于0，所以较大值减去较小值就可以完成
	// 不会存在tx=ty的情况
	for tx > sx && ty > sy && tx != ty {
		if tx > ty {
			tx = tx % ty
		} else {
			ty = ty % tx
		}
	}
	// 需要保证正数所有存在不够减的情况，
	switch {
	case tx == sx && ty == sy:
		return true
	case tx == sx: // 有一对x和y相等就会退出上面的循环，所以需要下面两个case处理
		return ty > sy && (ty-sy)%tx == 0
	case ty == sy:
		return tx > sx && (tx-sx)%ty == 0
	default:
		return false
	}
}
