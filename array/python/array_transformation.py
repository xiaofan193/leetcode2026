from typing import List

# 1243 数组的变换
#
# https://leetcode.cn/problems/array-transformation/description/
# 首先，给你一个初始数组 arr。然后，每天你都要根据前一天的数组生成一个新的数组。
#
# 第 i 天所生成的数组，是由你对第 i-1 天的数组进行如下操作所得的：
#
# 假如一个元素小于它的左右邻居，那么该元素自增 1。
# 假如一个元素大于它的左右邻居，那么该元素自减 1。
# 首、尾元素 永不 改变。
# 过些时日，你会发现数组将会不再发生变化，请返回最终所得到的数组。
#
# 例 1：
#
# 输入：[6,2,3,4]
# 输出：[6,3,3,4]
# 解释：
# 第一天，数组从 [6,2,3,4] 变为 [6,3,3,4]。
# 无法再对该数组进行更多操作。
# 示例 2：
#
# 输入：[1,6,3,4,3,5]
# 输出：[1,4,4,4,4,5]
# 解释：
# 第一天，数组从 [1,6,3,4,3,5] 变为 [1,5,4,3,4,5]。
# 第二天，数组从 [1,5,4,3,4,5] 变为 [1,4,4,4,4,5]。
# 无法再对该数组进行更多操作。


class Solution:
    def transformArray(self, arr: List[int]) -> List[int]:
        n = len(arr)
        # 首、尾元素永不改变，因此长度<=2时数组不会发生任何变化
        if n <= 2:
            return arr

        while True:
            changed = False
            # 1.每一天都基于前一天的数组整体生成新数组，所以先拷贝一份
            next_arr = arr[:]

            # 2.只需处理首尾之间的元素
            for i in range(1, n - 1):
                if arr[i] < arr[i - 1] and arr[i] < arr[i + 1]:
                    # 小于左右邻居，自增1
                    next_arr[i] += 1
                    changed = True
                elif arr[i] > arr[i - 1] and arr[i] > arr[i + 1]:
                    # 大于左右邻居，自减1
                    next_arr[i] -= 1
                    changed = True

            arr = next_arr
            # 3.某一天数组不再发生变化，即为最终结果
            if not changed:
                return arr
