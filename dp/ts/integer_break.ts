function integerBreak(n: number): number {
    /**
        dp[i]: i对应的最大乘积
        dp[2]: 1;
        ...
        dp[i]: max(
            1 * dp[i - 1], 1 * (i - 1),
            2 * dp[i - 2], 2 * (i - 2),
            ..., (i - 2) * dp[2], (i - 2) * 2
        );
     */
    const dp: number[] = new Array(n + 1).fill(0);
    dp[2] = 1;
    for (let i = 3; i <= n; i++) {
        for (let j = 1; j <= i / 2; j++) {
            dp[i] = Math.max(dp[i], j * dp[i - j], j * (i - j));
        }
    }
    return dp[n];
};