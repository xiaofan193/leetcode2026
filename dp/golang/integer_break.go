package golang

// 343 整数拆分
// https://leetcode.cn/problems/integer-break/description/

// 给定一个正整数 n ，将其拆分为 k 个 正整数 的和（ k >= 2 ），并使这些整数的乘积最大化。
// 返回 你可以获得的最大乘积 。
// 示例 1:
// 输入:
// n = 2
// 输出:
// 1
// 解释: 2 = 1 + 1, 1 × 1 = 1。
// 示例 2:
// 输入:
// n = 10
// 输出:
// 36
// 解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36。

func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 1

	for i := 3; i < n+1; i++ {
		for j := 1; j < i-1; j++ {
			dp[i] = max(dp[i], max(dp[i-j]*j, j*(i-j)))
		}

	}
	return dp[n]

}
