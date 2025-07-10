package medium

var graf399 map[string]map[string]float64
var visited399 map[string]struct{}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	graf399 = make(map[string]map[string]float64)
	// init graf399
	for i, v := range values {
		if graf399[equations[i][0]] == nil {
			graf399[equations[i][0]] = make(map[string]float64)
		}
		graf399[equations[i][0]][equations[i][1]] = v

		if graf399[equations[i][1]] == nil {
			graf399[equations[i][1]] = make(map[string]float64)
		}
		graf399[equations[i][1]][equations[i][0]] = 1.0 / v
	}

	result := make([]float64, 0, len(queries))
	for i, _ := range queries {
		root := queries[i][0]
		finish := queries[i][1]
		_, rootOk := graf399[root]
		_, finishOk := graf399[finish]
		if !rootOk || !finishOk {
			result = append(result, -1.0)
			continue
		}
		if root == finish {
			result = append(result, 1.0)
			continue
		}
		visited399 = make(map[string]struct{}, len(values)*2)
		result = append(result, dfs399(root, finish, 1.0))
	}
	return result
}

func dfs399(current, finish string, ans float64) float64 {
	if current == finish {
		return ans
	}
	if _, ok := visited399[current]; ok {
		return -1.0
	}

	visited399[current] = struct{}{}
	for k, v := range graf399[current] {
		temp := dfs399(k, finish, ans*v)
		if temp != -1.0 {
			return temp
		}
	}
	return -1.0
}
