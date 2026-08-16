function canPartition(nums: number[]): boolean {
    /**
        weightArr = nums;
        valueArr = nums;
        bagSize = sum / 2; (sum为nums各元素总和);
        按照0-1背包处理
     */
    const sum: number = nums.reduce((pre, cur) => pre + cur);
    if (sum % 2 === 1) return false;
    const bagSize: number = sum / 2;
    const weightArr: number[] = nums;
    const valueArr: number[] = nums;
    const goodsNum: number = weightArr.length;
    const dp: number[][] = new Array(goodsNum)
    	.fill(0)
    	.map(_ => new Array(bagSize + 1).fill(0));
    for (let i = weightArr[0]; i <= bagSize; i++) {
        dp[0][i] = valueArr[0];
    }
    for (let i = 1; i < goodsNum; i++) {
        for (let j = 0; j <= bagSize; j++) {
            if (j < weightArr[i]) {
                dp[i][j] = dp[i - 1][j];
            } else {
                dp[i][j] = Math.max(dp[i - 1][j], dp[i - 1][j - weightArr[i]] + valueArr[i]);
            }
        }
    }
    return dp[goodsNum - 1][bagSize] === bagSize;
};