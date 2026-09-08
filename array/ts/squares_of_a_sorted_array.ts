// 977. 有序数组的平方 (Squares of a Sorted Array)
// 双指针：平方后最大值只可能来自首尾两端，从结果数组末尾往前填。
function sortedSquares(nums: number[]): number[] {
    const n = nums.length;
    const res = new Array<number>(n);

    let i = 0, j = n - 1;
    for (let k = n - 1; k >= 0; k--) {
        const a = nums[i] * nums[i];
        const b = nums[j] * nums[j];
        if (a > b) {
            res[k] = a;
            i++;
        } else {
            res[k] = b;
            j--;
        }
    }
    return res;
}
