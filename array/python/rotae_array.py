class Solution:
    def rotate(self, nums: list[int], k: int) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        n = len(nums)
        if n==1: return
        k = k %  n
        self.reverse(nums,0,n-k-1)
        self.reverse(nums,n-k,n-1)
        self.reverse(nums,0,n-1)

    def reverse(self,nums,s,t):
        i,j = s,t
        while i < j:
            nums[i],nums[j] = nums[j],nums[i]
            i+=1
            j-=1