// 1049. 最后一块石头的重量 II
// 二维数组方式：dp[i][j] 表示在石头下标 [0, i] 之间任取，凑出的总重量不超过 j 的最大值
function lastStoneWeightII(stones: number[]): number {
    const sum: number = stones.reduce((a: number, b: number): number => a + b);
    const target: number = Math.floor(sum / 2);
    const n: number = stones.length;
    // dp[i][j]：容量（背包能装的重量上限）为 j，只考虑前 i+1 块石头（下标 [0, i]）能装下的最大重量
    const dp: number[][] = Array.from({ length: n }, () => new Array(target + 1).fill(0));

    // 初始化第 0 块石头：能放得下就放
    for (let j: number = stones[0]; j <= target; j++) {
        dp[0][j] = stones[0];
    }

    // 背包递推
    for (let i: number = 1; i < n; i++) {
        for (let j: number = 0; j <= target; j++) {
            if (stones[i] > j) {
                // 放不下第 i 块，只能不放
                dp[i][j] = dp[i - 1][j];
            } else {
                // 不放 vs 放：取较大者
                dp[i][j] = Math.max(dp[i - 1][j], dp[i - 1][j - stones[i]] + stones[i]);
            }
        }
    }

    // 分成两组抵消，差值 = sum - 2 * 尽量接近 sum/2 的那一组
    return sum - dp[n - 1][target] - dp[n - 1][target];
}