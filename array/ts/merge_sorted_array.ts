// 88 合并两个有序数组
// https://leetcode.cn/problems/merge-sorted-array/
// nums1 长度已是 m+n，末尾 n 个为占位 0，需原地修改 nums1（do not return anything, modify nums1 in-place instead）。
// 从后往前比较并填充，避免覆盖 nums1 前部未处理元素；j 清空即归并完成。
/**
 Do not return anything, modify nums1 in-place instead.
 */
function merge(nums1: number[], m: number, nums2: number[], n: number): void {
    let i = m - 1;
    let j = n - 1;
    for (let k = m + n - 1; j >= 0; k--) {
        if (i >= 0 && nums1[i] > nums2[j]) {
            nums1[k] = nums1[i];
            i--;
        } else {
            nums1[k] = nums2[j];
            j--;
        }
    }
}
