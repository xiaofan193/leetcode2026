# 1920. 基于排列构建数组 (Build Array from Permutation)
# https://leetcode.cn/problems/build-array-from-permutation/description/
# 思路：预分配等长数组，逐个按下标赋值 ans[i] = nums[nums[i]]，一次遍历 O(n)。
from typing import List


class Solution:
    def buildArray(self, nums: List[int]) -> List[int]:
        # 必须预分配成长度 len(nums)（或改用 append），空列表不能用下标赋值
        ans = [0] * len(nums)

        # range(len(nums)) 覆盖全部下标；写成 len(nums) - 1 会漏掉最后一个元素
        for i in range(len(nums)):
            ans[i] = nums[nums[i]]

        return ans
