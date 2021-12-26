package combination_sum

var (
	results     [][]int
	gTarget     int
	gCandidates []int
)

type Record struct {
	Val, Num int
}

func dfs(step, current int, buffer []*Record) {
	if current == gTarget {
		result := make([]int, 0)
		for _, b := range buffer {
			for i := 0; i < b.Num; i++ {
				result = append(result, b.Val)
			}
		}
		results = append(results, result)
		return
	}
	if step == len(gCandidates) || current > gTarget {
		return
	}
	for i := (gTarget - current) / gCandidates[step]; i >= 0; i-- {
		dfs(step+1, current+i*gCandidates[step], append(buffer, &Record{Val: gCandidates[step], Num: i}))
	}
}

func combinationSum(candidates []int, target int) [][]int {
	gTarget = target
	gCandidates = candidates
	results = make([][]int, 0)
	dfs(0, 0, []*Record{})
	return results
}
