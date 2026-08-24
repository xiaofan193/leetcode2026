package golang

// 494 目标和
// https://leetcode.cn/problems/target-sum/description/

// 给你一个非负整数数组 nums 和一个整数 target 。

// 向数组中的每个整数前添加 '+' 或 '-' ，然后串联起所有整数，可以构造一个 表达式 ：

// 例如，nums = [2, 1] ，可以在 2 之前添加 '+' ，在 1 之前添加 '-' ，然后串联起来得到表达式 "+2-1" 。
// 返回可以通过上述方法构造的、运算结果等于 target 的不同 表达式 的数目。

// 示例 1：

// 输入：nums = [1,1,1,1,1], target = 3
// 输出：5
// 解释：一共有 5 种方法让最终目标和为 3 。
// -1 + 1 + 1 + 1 + 1 = 3
// +1 - 1 + 1 + 1 + 1 = 3
// +1 + 1 - 1 + 1 + 1 = 3
// +1 + 1 + 1 - 1 + 1 = 3
// +1 + 1 + 1 + 1 - 1 = 3
// 示例 2：

// 输入：nums = [1], target = 1
// 输出：1

import "math"

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if math.Abs(float64(target)) > float64(sum) {
		return 0 // 此时没有方案
	}
	if (target+sum)%2 == 1 {
		return 0 // 此时没有方案
	}
	bagSize := (target + sum) / 2

	dp := make([][]int, len(nums))
	for i := range dp {
		dp[i] = make([]int, bagSize+1)
	}

	// 初始化最上行
	if nums[0] <= bagSize {
		dp[0][nums[0]] = 1
	}

	// 初始化最左列，最左列其他数值在递推公式中就完成了赋值
	dp[0][0] = 1

	var numZero float64
	for i := range nums {
		if nums[i] == 0 {
			numZero++
		}
		dp[i][0] = int(math.Pow(2, numZero))
	}

	// 以下遍历顺序行列可以颠倒
	for i := 1; i < len(nums); i++ { // 行，遍历物品
		for j := 0; j <= bagSize; j++ { // 列，遍历背包
			if nums[i] > j {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] + dp[i-1][j-nums[i]]
			}
		}
	}
	return dp[len(nums)-1][bagSize]
}
