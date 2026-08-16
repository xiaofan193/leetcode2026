package golang

// https://leetcode.cn/problems/last-stone-weight-ii/
// 1049. 最后一块石头的重量 II
// 有一堆石头，用整数数组 stones 表示。其中 stones[i] 表示第 i 块石头的重量。

// 每一回合，从中选出任意两块石头，然后将它们一起粉碎。假设石头的重量分别为 x 和 y，且 x <= y。那么粉碎的可能结果如下：

// 如果 x == y，那么两块石头都会被完全粉碎；
// 如果 x != y，那么重量为 x 的石头将会完全粉碎，而重量为 y 的石头新重量为 y-x。
// 最后，最多只会剩下一块 石头。返回此石头 最小的可能重量 。如果没有石头剩下，就返回 0。

func lastStoneWeightII(stones []int) int {
	// sum
	sum := 0

	for _, val := range stones {
		sum += val
	}
	target := sum / 2

	// 初始化
	dp := make([][]int, len(stones))

	for i := range dp {
		dp[i] = make([]int, target+1)
	}

	for j := stones[0]; j <= target; j++ {
		dp[0][j] = stones[0]
	}

	// 背包公式
	// dp[i][j]= max(dp[i-1][j],dp[i-1][j-stones[i]] + stones[j])
	for i := 1; i < len(stones); i++ {
		for j := 0; j <= target; j++ {
			if stones[i] > j {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-stones[i]]+stones[j])

			}
		}
	}
	return (sum - dp[len(stones)-1][target]) - dp[len(stones)-1][target]
}
