package golang
// min_cost_climbing_stairs.go
// https://leetcode.cn/problems/min-cost-climbing-stairs/
// 756. 使用最小花费爬楼梯
func minCostClimbingStairs(cost []int) int {
    if len(cost) <= 1{
        return 0
    }
   
    dp := make([]int,len(cost)+1)
    dp[0] = 0
    dp[1] = 0
   
    for i :=2; i <= len(cost);i++ {
        if dp[i-1] + cost[i-1] > dp[i-2]+ cost[i-2] {
            dp[i] = dp[i-2] + cost[i-2]
        }else{
            dp[i] = dp[i-1] + cost[i-1]
        }

    }
    return dp[len(cost)]
}