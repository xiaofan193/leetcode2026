# 977. 有序数组的平方 (Squares of a Sorted Array)
# 双指针：平方后最大值只可能来自首尾两端，从结果数组末尾往前填。
from typing import List

class Solution:
    def sortedSquares(self, nums: List[int]) -> List[int]:
        n = len(nums)
        res = [0] * n

        i, j = 0, n - 1
        for k in range(n - 1, -1, -1):
            a, b = nums[i] * nums[i], nums[j] * nums[j]
            if a > b:
                res[k] = a
                i += 1
            else:
                res[k] = b
                j -= 1

        return res
