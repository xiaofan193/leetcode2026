// 70 爬楼梯
// 二维数组方式：dp[i][j] 表示到达第 i 阶，且最后一步走了 j 阶（j = 1 或 2）的方法数
function climbStairs(n: number): number {
  if (n <= 1) {
    return n;
  }

  const dp: number[][] = Array.from({ length: n + 1 }, () => [0, 0, 0]); // 下标 1、2 分别代表最后一步走了 1、2 阶

  dp[1][1] = 1; // 从第 0 阶走 1 步到第 1 阶
  dp[2][1] = 1; // 从第 1 阶再走 1 步
  dp[2][2] = 1; // 从第 0 阶直接走 2 步

  for (let i = 3; i <= n; i++) {
    dp[i][1] = dp[i - 1][1] + dp[i - 1][2]; // 最后一步走 1 阶，前一状态在第 i-1 阶
    dp[i][2] = dp[i - 2][1] + dp[i - 2][2]; // 最后一步走 2 阶，前一状态在第 i-2 阶
  }

  return dp[n][1] + dp[n][2];
}