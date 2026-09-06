from typing import List


class Solution:
    # 88 合并两个有序数组
    # https://leetcode.cn/problems/merge-sorted-array/
    # nums1 长度已是 m+n，末尾 n 个为占位 0，需原地修改 nums1，不能新建数组返回。
    # 从后往前比较并填充，避免覆盖 nums1 前部未处理元素；j 清空即归并完成。
    def merge(self, nums1: List[int], m: int, nums2: List[int], n: int) -> None:
        i, j = m - 1, n - 1
        k = m + n - 1
        while j >= 0:
            if i >= 0 and nums1[i] > nums2[j]:
                nums1[k] = nums1[i]
                i -= 1
            else:
                nums1[k] = nums2[j]
                j -= 1
            k -= 1
