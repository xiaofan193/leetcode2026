function maxProfit(prices: number[]): number {
    if (prices.length === 0) return 0;
    let buy: number = prices[0];
    let profitMax: number = 0;
    for (let i = 1, length = prices.length; i < length; i++) {
        profitMax = Math.max(profitMax, prices[i] - buy);
        buy = Math.min(prices[i], buy);
    }
    return profitMax;
};

// DP
function maxProfit2(prices: number[]): number {
    /**
        dp[i][0]: 第i天持有股票的最大现金
        dp[i][1]: 第i天不持有股票的最大现金
     */
    const length = prices.length;
    if (length === 0) return 0;
    const dp: number[][] = [];
    dp[0] = [-prices[0], 0];
    for (let i = 1; i < length; i++) {
        dp[i] = [];
        dp[i][0] = Math.max(dp[i - 1][0], -prices[i]);
        dp[i][1] = Math.max(dp[i - 1][0] + prices[i], dp[i - 1][1]);
    }
    return dp[length - 1][1];
};