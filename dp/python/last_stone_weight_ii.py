from typing import List


class Solution:
    def lastStoneWeightII(self, stones: List[int]) -> int:
        total = sum(stones)
        target = total // 2
        n = len(stones)
        # 初始化 dp[i][j]：可以放 0-i 物品，背包容量为 j 的情况下背包中的最大价值
        dp = [[0] * (target + 1) for _ in range(n)]
        # dp[i][0] 默认初始化为 0
        # dp[0][j] 取决于 stones[0]
        for j in range(stones[0], target + 1):
            dp[0][j] = stones[0]

        for i in range(1, n):
            for j in range(1, target + 1):  # 注意是等于
                if j >= stones[i]:
                    # 不放: dp[i - 1][j]  放: dp[i - 1][j - stones[i]] + stones[i]
                    dp[i][j] = max(dp[i - 1][j], dp[i - 1][j - stones[i]] + stones[i])
                else:
                    dp[i][j] = dp[i - 1][j]

        return (total - dp[n - 1][target]) - dp[n - 1][target]
