package golang

// 0-1 背包问题（二维动态规划）
//
// 题目描述：有 n 件物品和一个容量为 bagWeight 的背包。
// 第 i 件物品的重量为 weight[i]，价值为 value[i]。
// 每件物品只能选择一次：装入背包（取 1）或不装（取 0），
// 求在不超过背包容量的前提下能获得的最大总价值。
//
// 状态定义：dp[i][j] 表示从下标 [0, i] 的物品中任意选取，
// 放入容量为 j 的背包所能达到的最大价值。
//
// 递推公式（对第 i 件物品做"选 / 不选"的决策）：
//   - 装不下（j < weight[i]）：dp[i][j] = dp[i-1][j]  // 只能不选，继承上一行的结果
//   - 装得下（j >= weight[i]）：dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i]]+value[i])
//     // 即"不选"与"选（腾出 weight[i] 空间，价值 +value[i]）"两者取较大
//
// 参考：https://programmercarl.com/algo/dynamic-programming/zero-one-knapsack-basics-part-1.html
func zeroOnePackage2D(weight, value []int, bagWeight int) int {
	n := len(weight)

	// dp[i][j]：容量为 j 时，从 [0, i] 的物品中取，最大价值
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, bagWeight+1)
	}

	// 初始化第一行（只考虑物品 0）
	// j < weight[0] 时容量装不下物品 0，保持默认值 0
	for j := weight[0]; j <= bagWeight; j++ {
		dp[0][j] = value[0]
	}

	for i := 1; i < n; i++ { // 遍历物品
		for j := 0; j <= bagWeight; j++ { // 遍历背包容量
			if j < weight[i] {
				dp[i][j] = dp[i-1][j] // 装不下，继承
			} else {
				// 取"不选"与"选"两者中的较大值
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-weight[i]]+value[i])
			}
		}
	}

	return dp[n-1][bagWeight]
}

// 0-1 背包（一维滚动数组优化）
//
// 观察二维递推式可知：dp[i][j] 只依赖 dp[i-1][...]（上一行），
// 因此可以把二维数组压缩成一维 dp[j]（容量为 j 的背包能装的最大价值）。
//
// 关键点：内层对容量 j 必须【倒序】遍历（j 从 bagWeight 递减到 weight[i]），
// 保证每个物品只被使用一次。如果正序遍历，dp[j-weight[i]] 可能已经是本轮
// 更新过的值，会导致同一件物品被重复放入，退化成完全背包问题。
//
// 参考：https://programmercarl.com/algo/dynamic-programming/zero-one-knapsack-basics-part-2.html
func zeroOnePackage1D(weight, value []int, bagWeight int) int {
	dp := make([]int, bagWeight+1) // 初始全为 0

	for i := 0; i < len(weight); i++ { // 遍历物品
		for j := bagWeight; j >= weight[i]; j-- { // 倒序遍历容量
			dp[j] = max(dp[j], dp[j-weight[i]]+value[i])
		}
	}

	return dp[bagWeight]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
