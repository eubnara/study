package leetcode

type Cache struct {
	r  []bool
	c  []bool
	dx []bool
	dy []bool
}

func check(x, y, n int, cache *Cache) {
	// row
	cache.r[x] = true
	// col
	cache.c[y] = true
	// dx
	cache.dx[n-1-x+y] = true
	// dy
	cache.dy[x+y] = true
}

func visited(x, y, n int, cache *Cache) bool {
	return cache.r[x] || cache.c[y] || cache.dx[n-1-x+y] || cache.dy[x+y]
}

func solveNQueens(n int) [][]string {
	cache := Cache{
		r:  make([]bool, n),
		c:  make([]bool, n),
		dx: make([]bool, 2*n-1),
		dy: make([]bool, 2*n-1),
	}
	return solveNQueen(n, 0, cache)
}

func makeQStr(l, t int) string {
	chars := []rune{}
	for i := 0; i < l; i++ {
		c := '.'
		if t == i {
			c = 'Q'
		}
		chars = append(chars, c)
	}
	return string(chars)
}

func solveNQueen(n int, cur int, cache Cache) [][]string {
	ans := [][]string{}
	for i := 0; i < n; i++ {
		if visited(cur, i, n, &cache) {
			continue
		}
		curStrArr := []string{makeQStr(n, i)}
		if cur == n-1 {
			ans = append(ans, curStrArr)
			continue
		}
		newCache := Cache{
			r:  make([]bool, n),
			c:  make([]bool, n),
			dx: make([]bool, 2*n-1),
			dy: make([]bool, 2*n-1),
		}
		copy(newCache.c, cache.c)
		copy(newCache.r, cache.r)
		copy(newCache.dx, cache.dx)
		copy(newCache.dy, cache.dy)
		check(cur, i, n, &newCache)
		for _, v := range solveNQueen(n, cur+1, newCache) {
			ans = append(ans, append(curStrArr, v...))
		}
	}

	return ans
}
