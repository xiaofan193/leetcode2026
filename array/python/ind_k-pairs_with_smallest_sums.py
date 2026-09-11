import heapq
from typing import List


class Solution:
    # 思路：把数对看成 m×n 的矩阵，行列都递增，问题变成"在有序矩阵里取前 k 小"。
    # 用堆维护"已发现的候选边界"：初始只有 (0,0)，每次弹出堆顶后再把它的右邻居和
    # 下邻居推进堆。为避免同一个格子被两条路径重复入堆，只让第一列向下扩展。
    def kSmallestPairs(self, nums1: List[int], nums2: List[int], k: int) -> List[List[int]]:
        # sorted 返回新列表，不改动调用方传入的数组
        a, b = sorted(nums1), sorted(nums2)

        ans = []
        if not a or not b or k <= 0:
            return ans

        # 堆元素是 (和, i, j)，元组按字典序比较，和相同时自然按 i、j 稳定排序
        heap = [(a[0] + b[0], 0, 0)]

        while len(ans) < k and heap:
            _, i, j = heapq.heappop(heap)
            ans.append([a[i], b[j]])

            # 同一行向右扩展
            if j + 1 < len(b):
                heapq.heappush(heap, (a[i] + b[j + 1], i, j + 1))

            # 只有第一列才向下扩展；否则 (i,j) 会被 (i-1,j) 和 (i,j-1) 重复入堆
            if j == 0 and i + 1 < len(a):
                heapq.heappush(heap, (a[i + 1] + b[0], i + 1, 0))

        # 对数不足 k 时堆会先空，这里自然返回全部，不会越界
        return ans
