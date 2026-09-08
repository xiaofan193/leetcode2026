class Solution:
    def intersection(self, nums1: List[int], nums2: List[int]) -> List[int]:
        res = []
        # 双指针前提：两个数组都要先排序
        nums1.sort()
        nums2.sort()
        m, n = len(nums1), len(nums2)

        i, j = 0, 0

        while i < m and j < n:
            if nums1[i] < nums2[j]:
                i += 1
            elif nums1[i] > nums2[j]:
                j += 1
            else:
                # 去重：和 res 末位相同则不再追加（排序后重复值必然连续）
                if not res or res[-1] != nums1[i]:
                    res.append(nums1[i])
                # 无论是否追加，指针都必须前进，否则死循环
                i += 1
                j += 1

        return res
