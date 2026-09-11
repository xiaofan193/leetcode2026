// 280 摆动排序
// https://leetcode.cn/problems/wiggle-sort/description/

// 给你一个没有排序的数组，请将原数组就地重新排列满足如下性质 nums[0] <= nums[1] >= nums[2] <= nums[3].... 请就地排序数组，也就是不需要额外数组
// For example, given nums = [3, 5, 2, 1, 6, 4], one possible answer is [1, 6, 2, 5, 3, 4]

// 一趟遍历就地调整：奇数位要"峰"(nums[i] >= nums[i-1])，偶数位要"谷"(nums[i] <= nums[i-1])。
// 不满足就交换相邻两位。i=0 时 nums[-1] 为 undefined，两个比较都是 false，不会误交换。
function wiggleSort(nums: number[]): number[] {
    for (let i = 0; i < nums.length; i++) {
        if (i % 2 === 1) {
            if (nums[i] < nums[i - 1]) {
                [nums[i], nums[i - 1]] = [nums[i - 1], nums[i]];
            }
        } else {
            if (nums[i] > nums[i - 1]) {
                [nums[i], nums[i - 1]] = [nums[i - 1], nums[i]];
            }
        }
    }
    return nums;
}
