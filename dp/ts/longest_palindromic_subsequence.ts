
function longestPalindromeSubseq(s: string): number {
    
    const size: number = s.length;
    
    const dp: number[][] = new Array(size).fill(0).map(_ => new Array(size).fill(0));

    for (let i = 0; i < size;i++) {
        dp[i][i] = 1;
    }

    for (let i = size - 1; i >= 0; i--) {
        for (let j = i + 1; j < size; j++) {
            if (s[i] === s[j]) {
                dp[i][j] = dp[i + 1][j - 1] + 2;
            } else {
                dp[i][j] = Math.max(dp[i + 1][j], dp[i][j - 1]);
            }
        }
    }
    return dp[0][size - 1];
}