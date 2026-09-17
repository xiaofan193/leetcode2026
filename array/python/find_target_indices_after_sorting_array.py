# 2089. 找出数组排序后的目标下标 (Find Target Indices After Sorting Array)
# https://leetcode.cn/problems/find-target-indices-after-sorting-array/description/
# 思路：不需要真的排序。排序后所有 target 会连续排在一起，起始下标只取决于"比 target 小的
# 元素有多少个"：设 less 为小于 target 的元素个数、equal 为等于 target 的个数，答案就是
# [less, less + equal)。一次遍历统计即可，时间 O(n)、额外空间 O(1)（不计结果数组）。
from typing import List


class Solution:
    def targetIndices(self, nums: List[int], target: int) -> List[int]:
        less, equal = 0, 0

        # 遍历的是元素值本身，不是下标；写成 i < target 会变成拿下标和 target 比
        for num in nums:
            if num < target:
                less += 1
            elif num == target:
                equal += 1

        ans = [0] * equal

        # 也可以直接写 list(range(less, less + equal))
        for i in range(len(ans)):
            ans[i] = less + i

        return ans
