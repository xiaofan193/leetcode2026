import heapq
from typing import List

# find_k-pairs_with_smallest_sums
# 373. 查找和最小的 K 对数字
# https://leetcode.cn/problems/find-k-pairs-with-smallest-sums/description/
#
# 给定两个以 非递减顺序排列 的整数数组 nums1 和 nums2 , 以及一个整数 k 。
#
# 定义一对值 (u,v)，其中第一个元素来自 nums1，第二个元素来自 nums2 。
#
# 请找到和最小的 k 个数对 (u1,v1),  (u2,v2)  ...  (uk,vk) 。
#
# 示例 1:
#
# 输入: nums1 = [1,7,11], nums2 = [2,4,6], k = 3
# 输出: [[1,2],[1,4],[1,6]]
# 解释: 返回序列中的前 3 对数：
#      [1,2],[1,4],[1,6],[7,2],[7,4],[11,2],[7,6],[11,4],[11,6]
# 示例 2:
#
# 输入: nums1 = [1,1,2], nums2 = [1,2,3], k = 2
# 输出: [[1,1],[1,1]]
# 解释: 返回序列中的前 2 对数：
#      [1,1],[1,1],[1,2],[2,1],[1,2],[2,2],[1,3],[1,3],[2,3]


class Solution:
    # 思路：把数对看成 m×n 的矩阵，行列都递增，问题变成"在有序矩阵里取前 k 小"。
    # 用堆维护"已发现的候选边界"：初始只有 (0,0)，每次弹出堆顶后再把它的右邻居和
    # 下邻居推进堆。为避免同一个格子被两条路径重复入堆，只让第一列向下扩展。
    def kSmallestPairs(self, nums1: List[int], nums2: List[int], k: int) -> List[List[int]]:
        # sorted 返回新列表，不改动调用方传入的数组
        nums1 = sorted(nums1)
        nums2 = sorted(nums2)

        ans: List[List[int]] = []
        if not nums1 or not nums2 or k <= 0:
            return ans

        # 堆元素是 (和, i, j)，元组按字典序比较，和相同时自然按 i、j 稳定排序
        heap = [(nums1[0] + nums2[0], 0, 0)]

        while len(ans) < k and heap:
            _, i, j = heapq.heappop(heap)
            ans.append([nums1[i], nums2[j]])

            # 同一行向右扩展
            if j + 1 < len(nums2):
                heapq.heappush(heap, (nums1[i] + nums2[j + 1], i, j + 1))

            # 只有第一列才向下扩展；否则 (i,j) 会被 (i-1,j) 和 (i,j-1) 重复入堆
            if j == 0 and i + 1 < len(nums1):
                heapq.heappush(heap, (nums1[i + 1] + nums2[0], i + 1, 0))

        # 对数不足 k 时堆会先空，这里自然返回全部，不会越界
        return ans
