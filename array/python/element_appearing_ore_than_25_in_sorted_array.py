from typing import List
class Solution:
    def findSpecialInteger(self, arr: List[int]) -> int:
        n = len(arr)
        c,k = arr[0],0

        for i in range(0,n):
            if arr[i] == c:
                k+=1
                if k*4 > n: return c
            else:
                c,k = arr[i],1

        return c