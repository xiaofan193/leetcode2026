// 1920 基于排列构建数组
// https://leetcode.cn/problems/build-array-from-permutation/
// 预分配等长数组，按下标依次填 ans[i] = nums[nums[i]]，一次遍历 O(n)。
function buildArray(nums: number[]): number[] {
    const n = nums.length;
    // 预分配成长度 n（等价于 Go 的 make([]int, len(nums))）
    const ans: number[] = new Array(n);

    // 取 nums[i] 作为下标再去 nums 里取值，注意是 nums[nums[i]] 不是 nums[i]
    for (let i = 0; i < n; i++) {
        ans[i] = nums[nums[i]];
    }

    // 少写这行 return 会返回 undefined，tsc 在 --strict 下也会直接报 TS2355
    return ans;
}
