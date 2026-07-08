package backtracking

//https://leetcode.cn/problems/non-decreasing-subsequences/description/
//491. 非递减子序列

// 给你一个整数数组 nums ，找出并返回所有该数组中不同的递增子序列，递增子序列中 至少有两个元素 。你可以按 任意顺序 返回答案。

// 数组中可能含有重复元素，如出现两个整数相等，也可以视作递增序列的一种特殊情况。

// 示例 1：

// 输入：nums = [4,6,7,7]
// 输出：[[4,6],[4,6,7],[4,6,7,7],[4,7],[4,7,7],[6,7],[6,7,7],[7,7]]
// 示例 2：

// 输入：nums = [4,4,3,2,1]
// 输出：[[4,4]]

func findSubsequences(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0, len(nums))

	var dfs func(start int)
	dfs = func(start int) {
		if len(path) >= 2 { // 空数组 和原来的数组接入结果
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
		}

		used := make(map[int]bool)
		for i := start; i < len(nums); i++ {
			if used[nums[i]] {
				continue
			}
			if len(path) > 0 && nums[i] < path[len(path)-1] {
				continue
			}

			used[nums[i]] = true
			path = append(path, nums[i])
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}

	dfs(0)
	return res
}
