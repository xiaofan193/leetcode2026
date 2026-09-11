package golang

//264. 丑数 II
//  https://leetcode.cn/problems/ugly-number-ii/description/

// 给你一个整数 n ，请你找出并返回第 n 个 丑数 。

// 丑数 就是质因子只包含 2、3 和 5 的正整数。

// 示例 1：

// 输入：n = 10
// 输出：12
// 解释：[1, 2, 3, 4, 5, 6, 8, 9, 10, 12] 是由前 10 个丑数组成的序列。
// 示例 2：

// 输入：n = 1
// 输出：1
// 解释：1 通常被视为丑数。

// 提示：

// 1 <= n <= 1690

// 三指针 DP：任何一个丑数都能由比它小的某个丑数乘 2/3/5 得到（1 除外）
// dp[i] = min(dp[p2]*2, dp[p3]*3, dp[p5]*5)，取到最小值的指针（可能多个）向前推进一步
func nthUglyNumber(n int) int {
	dp := make([]int, n)
	dp[0] = 1

	// p2、p3、p5 分别指向"下一个待乘以 2/3/5 的丑数"的下标
	p2, p3, p5 := 0, 0, 0

	for i := 1; i < n; i++ {
		a, b, c := dp[p2]*2, dp[p3]*3, dp[p5]*5

		min := a
		if b < min {
			min = b
		}
		if c < min {
			min = c
		}
		dp[i] = min

		// 必须用三个独立 if：当最小值同时由多个指针产出时（如 6=2*3=3*2），
		// 命中的指针都要前进，否则下一轮会重复写入同一个值
		if min == a {
			p2++
		}
		if min == b {
			p3++
		}
		if min == c {
			p5++
		}
	}

	return dp[n-1]
}
