package backtracking

// https://leetcode.cn/problems/combination-sum-iii/description/
// 216. 组合总和 III
// 找出所有相加之和为 n 的 k 个数的组合，且满足下列条件：
// 只使用数字 1 到 9，每个数字最多使用一次。返回所有可能的有效组合的列表。
//
// 示例 1：
// 输入: k = 3, n = 7
// 输出: [[1,2,4]]
// 示例 2：
// 输入: k = 3, n = 9
// 输出: [[1,2,6],[1,3,5],[2,3,4]]

var (
	path []int
	res  [][]int
)

func combinationSum3(k int, n int) [][]int {
	res, path = make([][]int, 0), make([]int, 0, k)
	backtracking(k, n, 1, 0)
	return res
}

func backtracking(k int, n int, start int, sum int) {
	if len(path) == k {
		if sum == n {
			tmp := make([]int, k)
			copy(tmp, path)
			res = append(res, tmp)
		}
		return
	}

	for i := start; i <= 9; i++ {
		// 剪枝：数字和已超过 n，或剩余可用的数字不足
		if sum+i > n || 9-i+1 < k-len(path) {
			break
		}
		path = append(path, i)
		backtracking(k, n, i+1, sum+i)
		path = path[:len(path)-1]
	}
}
