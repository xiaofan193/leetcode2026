from typing import List

class Solution:
    def minimumAbsDifference(self, arr: List[int]) -> List[List[int]]:
        res = []
        if not arr:
            return res

        arr.sort()
        min_diff = arr[1] - arr[0]

        for i in range(1, len(arr)):
            diff = arr[i] - arr[i - 1]
            item = [arr[i - 1], arr[i]]

            if diff == min_diff:
                res.append(item)

            if diff < min_diff:
                min_diff = diff
                res.clear()
                res.append(item)

        return res
