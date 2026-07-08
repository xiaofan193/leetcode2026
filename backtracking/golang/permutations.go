package backtracking

// https://leetcode.cn/problems/permutations/
// 46. 全排列
// 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。

// 示例 1：

// 输入：nums = [1,2,3]
// 输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
// 示例 2：

// 输入：nums = [0,1]
// 输出：[[0,1],[1,0]]
// 示例 3：

// 输入：nums = [1]
// 输出：[[1]]

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0, len(nums))
	used := make([]bool, len(nums))

	var dfs func()
	dfs = func() {
		if len(path) == len(nums) {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			path = append(path, nums[i])
			used[i] = true
			dfs()
			used[i] = false
			path = path[:len(path)-1]
		}
	}

	dfs()
	return res
}


var (
    res [][]int
    path  []int
    st    []bool   // abbreviation for state
)
func permute(nums []int) [][]int {
    res, path = make([][]int, 0), make([]int, 0, len(nums))
    st = make([]bool, len(nums))
    dfs_permute(nums, 0)
    return res
}

func dfs_permute(nums []int, cur int) {
    if cur == len(nums) {
        tmp := make([]int, len(path))
        copy(tmp, path)
        res = append(res, tmp)
    }
    for i := 0; i < len(nums); i++ {
        if !st[i] {
            path = append(path, nums[i])
            st[i] = true
            dfs(nums, cur + 1)
            st[i] = false
            path = path[:len(path)-1]
        }
    }
}

// https://leetcode.com/problems/permutations-ii/description/

func permuteUnique(nums []int) [][]int {
    res, path = make([][]int, 0), make([]int, 0, len(nums))
    st = make([]bool, len(nums))
    sort.Ints(nums)
    dfs_permuteUnique(nums, 0)
    return res
}

func dfs_permuteUnique(nums []int, cur int) {
    if cur == len(nums) {
        tmp := make([]int, len(path))
        copy(tmp, path)
        res = append(res, tmp)
    }
    for i := 0; i < len(nums); i++ {
        if i != 0 && nums[i] == nums[i-1] && !st[i-1] {  // Deduplication, using st to differentiate levels
            continue
        }
        if !st[i] {
            path = append(path, nums[i])
            st[i] = true
            dfs(nums, cur + 1)
            st[i] = false
            path = path[:len(path)-1]
        }
    }