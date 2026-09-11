from typing import List


class Solution:
    # 360 有序数组转化
    # https://leetcode.cn/problems/sort-transformed-array/
    # 给定升序数组 nums 和 a、b、c，对每个 x 应用 f(x) = ax^2 + bx + c，返回升序结果。

    # 方法一：双指针 O(n)
    # 抛物线 f(x) 是对称的，在升序的 nums 上取值先降后升（a>0）或先升后降（a<0），
    # 也就是最大/最小值一定出现在两端。于是从两端往中间收拢，
    # 每次取"更大"的那个填到结果末尾（开口向上），或取"更小"的填到结果开头（开口向下）。
    # nums[i-1] 这类越界在这里不存在：两侧各自独立推进，指针不会交叉。
    def sortTransformedArray(self, nums: List[int], a: int, b: int, c: int) -> List[int]:
        def f(x: int) -> int:
            return a * x * x + b * x + c

        n = len(nums)
        res = [0] * n
        left, right = 0, n - 1
        # a >= 0 时从末尾往前填（先放大的）；a < 0 时从头往后填（先放小的）
        k = n - 1 if a >= 0 else 0

        while left <= right:
            left_val, right_val = f(nums[left]), f(nums[right])
            if a >= 0:
                if left_val >= right_val:
                    res[k] = left_val
                    left += 1
                else:
                    res[k] = right_val
                    right -= 1
                k -= 1
            else:
                if left_val <= right_val:
                    res[k] = left_val
                    left += 1
                else:
                    res[k] = right_val
                    right -= 1
                k += 1

        return res

    # 方法二：朴素解法 O(n log n)
    # 逐个算出 f(x) 再整体排序，不利用 nums 有序的性质，用来对拍方法一。
    def sortTransformedArrayBruteForce(
        self, nums: List[int], a: int, b: int, c: int
    ) -> List[int]:
        return sorted(a * x * x + b * x + c for x in nums)
