// 360 有序数组转化
// https://leetcode.cn/problems/sort-transformed-array/
// 给定升序数组 nums 和 a、b、c，对每个 x 应用 f(x) = ax^2 + bx + c，返回升序结果。

// 方法一：双指针 O(n)
// 抛物线 f(x) 是对称的，在升序的 nums 上取值先降后升（a>0）或先升后降（a<0），
// 也就是最大/最小值一定出现在两端。于是从两端往中间收拢，
// 每次取"更大"的那个填到结果末尾（开口向上），或取"更小"的填到结果开头（开口向下）。
// 循环条件是 left <= right，否则 left === right 的那个元素会被漏掉。
function sortTransformedArray(nums: number[], a: number, b: number, c: number): number[] {
    const f = (x: number): number => a * x * x + b * x + c;

    const n = nums.length;
    const res = new Array<number>(n);
    let left = 0;
    let right = n - 1;
    // a >= 0 时从末尾往前填（先放大的）；a < 0 时从头往后填（先放小的）
    let k = a >= 0 ? n - 1 : 0;

    while (left <= right) {
        const leftVal = f(nums[left]);
        const rightVal = f(nums[right]);
        if (a >= 0) {
            if (leftVal >= rightVal) {
                res[k] = leftVal;
                left++;
            } else {
                res[k] = rightVal;
                right--;
            }
            k--;
        } else {
            if (leftVal <= rightVal) {
                res[k] = leftVal;
                left++;
            } else {
                res[k] = rightVal;
                right--;
            }
            k++;
        }
    }

    return res;
}

// 方法二：朴素解法 O(n log n)
// 逐个算出 f(x) 再整体排序，不利用 nums 有序的性质，用来对拍方法一。
// 注意必须传比较函数，否则 JS 的默认排序是按字符串字典序。
function sortTransformedArrayBruteForce(nums: number[], a: number, b: number, c: number): number[] {
    return nums.map((x) => a * x * x + b * x + c).sort((p, q) => p - q);
}
