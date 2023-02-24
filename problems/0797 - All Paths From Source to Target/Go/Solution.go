// Runtime: 12 ms
// Memory Usage: 7.8 MB

func allPathsSourceTarget(graph [][]int) [][]int {
	var res [][]int
	var cur []int
	cur = append(cur, 0)
	dfs(graph, 0, cur, &res)

	return res
}

func dfs(graph [][]int, cur int, curPath []int, res *[][]int) {
	if cur == len(graph)-1 {
		tmp := make([]int, len(curPath))
		copy(tmp, curPath)
		*res = append(*res, tmp)
		return
	}

	for _, v := range graph[cur] {
		tmp := make([]int, len(curPath))
		copy(tmp, curPath)
		tmp = append(tmp, v)
		dfs(graph, v, tmp, res)
	}
}