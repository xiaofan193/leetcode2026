package golang

import "math"

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/description/
// 121 买卖股票的最佳时机

// 给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。

// 你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。

// 返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。

// 示例 1：

// 输入：[7,1,5,3,6,4]
// 输出：5
// 解释：在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
//      注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格；同时，你不能在买入前卖出股票。
// 示例 2：

// 输入：prices = [7,6,4,3,1]
// 输出：0
// 解释：在这种情况下, 没有交易完成, 所以最大利润为 0

// 先用弹性算法比较简单
func maxProfit(prices []int) int {
	cost := math.MaxInt32
	profit := 0

	for _, price := range prices {
		cost = min(cost, price)
		profit = max(profit, price-cost)
	}
	return profit
}

func maxProfit2(prices []int) int {
	min := prices[0]
	res := 0
	for i := 1; i < len(prices); i++ {
		if prices[i]-min > res {
			res = prices[i] - min
		}
		if min > prices[i] {
			min = prices[i]
		}
	}
	return res
}

// 动态规划
// dp[i][0] = max(dp[i - 1][0], -prices[i]); 和 dp[i][1] = max(dp[i - 1][1], prices[i] + dp[i - 1][0]);
// dp[i][0] 表示第i天持有股票所得最多现金
// dp[i][1] 表示第i天不持有股票所得最多现金

func maxProfit3(prices []int) int {
	length := len(prices)
	if length == 0 {
		return 0
	}
	dp := make([][]int, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	for i := 1; i < length; i++ {
		dp[i][0] = max(dp[i-1][0], -prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
	}
	return dp[length-1][1]
}
