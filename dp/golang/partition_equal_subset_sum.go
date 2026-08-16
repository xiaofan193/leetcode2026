package golang

// 416 . 分割等和子集
// https://leetcode.cn/problems/partition-equal-subset-sum/description/

// 给你一个 只包含正整数 的 非空 数组 nums 。请你判断是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。

// 示例 1：

// 输入：nums = [1,5,11,5]
// 输出：true
// 解释：数组可以分割成 [1, 5, 5] 和 [11] 。
// 示例 2：

// 输入：nums = [1,2,3,5]
// 输出：false
// 解释：数组不能分割成两个元素和相等的子集。

func canPartition(nums []int) bool {
	sum := 0
	for _, val := range nums {
		sum += val
	}
	if sum%2 == 1 {
		return false
	}
	target := sum / 2
	dp := make([][]int, len(nums))
	for i := range dp {
		dp[i] = make([]int, target+1)
	}
	for j := nums[0]; j <= target; j++ {
		dp[0][j] = nums[0]
	}
	for i := 1; i < len(nums); i++ {
		for j := 0; j <= target; j++ {
			if j < nums[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-nums[i]]+nums[i])
			}
		}
	}
	return dp[len(nums)-1][target] == target
}
