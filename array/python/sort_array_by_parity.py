from typing import List


class Solution:
    # 905 按奇偶排序数组
    # https://leetcode.cn/problems/sort-array-by-parity/
    # 把偶数移到前面、奇数移到后面，不要求保持原相对顺序，就地修改即可。

    # 双指针：i 从左找"最左边的奇数"，j 从右找"最右边的偶数"，
    # 两个都找到就交换（一个奇数换到右边、一个偶数换到左边），然后问题缩小为 [i+1, j-1] 的子问题。
    # 注意判断奇数用 `% 2 != 0` 而不是 `== 1`：Python 里 -3 % 2 == 1 虽然也对，
    # 但写成 != 0 与 TS/JS 的行为一致，负数不会踩坑。
    def sortArrayByParity(self, nums: List[int]) -> List[int]:
        i, j = 0, len(nums) - 1
        while i < j:  # 循环直到不足两个数
            if nums[i] % 2 == 0:  # 当前是偶数，已经在正确一侧
                i += 1
            elif nums[j] % 2 != 0:  # 当前是奇数，已经在正确一侧
                j -= 1
            else:  # nums[i] 是奇数、nums[j] 是偶数，交换
                nums[i], nums[j] = nums[j], nums[i]
                i += 1
                j -= 1
        return nums
