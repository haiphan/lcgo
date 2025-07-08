package problems

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	N := len(equations)
	dist := make(map[string]map[string]float64)
	setDist := func(a, b string, x float64) bool {
		_, hasA := dist[a]
		if !hasA {
			dist[a] = make(map[string]float64)
			dist[a][b] = x
			return true
		} else {
			_, hasB := dist[a][b]
			if !hasB {
				dist[a][b] = x
			}
			return !hasB
		}
	}
	for i, e := range equations {
		a, b := e[0], e[1]
		v := values[i]
		setDist(a, a, 1)
		setDist(b, b, 1)
		setDist(a, b, v)
		setDist(b, a, 1/v)
	}
	for r := 0; r < N; r++ {
		mod := false
		for i, e := range equations {
			a, b := e[0], e[1]
			v := values[i]
			for c, cd := range dist[b] {
				nd := v * cd
				mod = mod || setDist(a, c, nd)
				setDist(c, a, 1/nd)
			}
			for c, cd := range dist[a] {
				nd := cd / v
				mod = mod || setDist(b, c, nd)
				setDist(c, b, 1/nd)
			}
		}
		if !mod {
			break
		}
	}
	res := make([]float64, len(queries))
	for i := 0; i < len(queries); i++ {
		a, b := queries[i][0], queries[i][1]
		r := -1.0
		_, hasA := dist[a]
		if hasA {
			d, hasB := dist[a][b]
			if hasB {
				r = d
			}
		}
		res[i] = r
	}
	return res
}
