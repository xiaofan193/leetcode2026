package backtracking
// https://leetcode.cn/problems/combination-sum/description/
// 给你一个 无重复元素 的整数数组 candidates 和一个目标整数 target ，找出 candidates 中可以使数字和为目标数 target 的 所有 不同组合 ，并以列表形式返回。你可以按 任意顺序 返回这些组合。
// candidates 中的 同一个 数字可以 无限制重复被选取 。如果至少一个数字的被选数量不同，则两种组合是不同的。 
// 对于给定的输入，保证和为 target 的不同组合数少于 150 个。
// 输入：candidates = [2,3,6,7], target = 7
// 输出：[[2,2,3],[7]]
// 解释：
// 2 和 3 可以形成一组候选，2 + 2 + 3 = 7 。注意 2 可以使用多次。
// 7 也是一个候选， 7 = 7 。
// 仅有这两种组合
var (
    res [][]int
    path  []int
)
func combinationSum(candidates []int, target int) [][]int {
    res, path = make([][]int, 0), make([]int, 0, len(candidates))
    sort.Ints(candidates)   // Sort for pruning
    dfs(candidates, 0, target)
    return res
}

func dfs(candidates []int, start int, target int) {
    if target == 0 {
        tmp := make([]int, len(path))
        copy(tmp, path)
        res = append(res, tmp)
        return
    }
    for i := start; i < len(candidates); i++ {
        if candidates[i] > target {
            break
        }
        path = append(path, candidates[i])
        dfs(candidates, i, target - candidates[i])
        path = path[:len(path) - 1]
    }
}
// https://leetcode.cn/problems/combination-sum-ii/description/

// 给定一个候选人编号的集合 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

// candidates 中的每个数字在每个组合中只能使用 一次 。

// 注意：解集不能包含重复的组合。 

 

// 示例 1:

// 输入: candidates = [10,1,2,7,6,1,5], target = 8,
// 输出:
// [
// [1,1,6],
// [1,2,5],
// [1,7],
// [2,6]
// ]
// 示例 2:

// 输入: candidates = [2,5,2,1,2], target = 5,
// 输出:
// [
// [1,2,2],
// [5]
// ]

func combinationSum2(candidates []int,target int) {
	res = make([][]int,0)
	path = make([]int,0,len(candidates))
	sort.Ints(candidates)
	dfs2(candidates,0,target)
	return res 
}  


func dfs2(candidates []int,start int,target int){
	if target == 0 {
        temp := make([]int,0)
		copy(temp,path)
		res = append(res,tmp)
		return
	}

	for i := start;i <len(candidates);i++{
		if candidates[i] > target {
			break 
		}

		if i != start && candidates[i] == candidates[i-1] {
			 continue 
		}

		path = append(path,candidates[i])
		dfs2(candidates,i + 1,target - candidates[i])

		path = path[:len(path) - 1]
	}

}