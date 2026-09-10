from typing import List

# 1213 3个有序数组的交集
# intersection_of_three_sorted_arrays
#
# https://leetcode.cn/problems/intersection-of-three-sorted-arrays/description/
#
# 给定3个严格递增的排列的整数数组 arr1 arr2 arr3,返回一个仅由这3个数组中同时出现整数构成的有序数组
# 例如： arr1 = [1,2,3,4,5], arr2 = [1,2,5,7,9],arr3 = [1,3,4,5,8]
# 答案: [1,5]

class Solution:
    def arraysIntersection(self, arr1: List[int], arr2: List[int], arr3: List[int]) -> List[int]:
        ans = []
        # 三指针同向扫描：三个指针都指向各自数组当前最小的未比较元素
        i, j, k = 0, 0, 0

        while i < len(arr1) and j < len(arr2) and k < len(arr3):
            if arr1[i] == arr2[j] and arr2[j] == arr3[k]:
                # 数组严格递增，相等元素唯一，无需去重
                ans.append(arr1[i])
                i += 1
                j += 1
                k += 1
            else:
                # 把三者中最小的指针右移，向更大的值靠拢
                minx = arr1[i]
                if arr2[j] < minx:
                    minx = arr2[j]
                if arr3[k] < minx:
                    minx = arr3[k]

                # 可能多个指针同时命中最小值，都要前进
                if arr1[i] == minx:
                    i += 1

                if arr2[j] == minx:
                    j += 1

                if arr3[k] == minx:
                    k += 1

        return ans
