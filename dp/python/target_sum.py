class Solution:
    def findTargetSumWays(self, nums: List[int], target: int) -> int:
        total_sum = sum(nums)  # 计算nums的总和
        if abs(target) > total_sum:
            return 0  # 此时没有方案
        if (target + total_sum) % 2 == 1:
            return 0  # 此时没有方案
        target_sum = (target + total_sum) // 2  # 目标和

        # 创建二维动态规划数组，行表示选取的元素数量，列表示累加和
        dp = [[0] * (target_sum + 1) for _ in range(len(nums) + 1)]
        dp = [[0] * (target_sum + 1) for _ in range(len(nums))]

        # 初始化状态
        dp[0][0] = 1
        if nums[0] <= target_sum:
            dp[0][nums[0]] = 1
        numZero = 0
        for i in range(len(nums)):
            if nums[i] == 0:
                numZero += 1
            dp[i][0] = int(math.pow(2, numZero))

        # 动态规划过程
        for i in range(1, len(nums)):
            for j in range(target_sum + 1):
                dp[i][j] = dp[i - 1][j]  # 不选取当前元素
                if j >= nums[i - 1]:
                    dp[i][j] += dp[i - 1][j - nums[i]]  # 选取当前元素

        return dp[len(nums)-1][target_sum]  # 返回达到目标和的方案数