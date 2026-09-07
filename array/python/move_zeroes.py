class Solution:
    def moveZeroes(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        n = len(nums)
        k = 0
        for i in range (0,n):
            if i != 0:
                nums[k] = nums[i]
                k+= 1


        while (k < n):
            nums[k] = 0
            k+= 1