class Solution:
    def removeDuplicates2(self, nums: List[int]) -> int:
        n = len(nums)
        if n == 1:
            return 1
        k = 2
        for i in range(2, n):
            if not ((nums[k - 2] == nums[k - 1]) and (nums[i] == nums[k - 1])):
                nums[k] = nums[i]
                k += 1
        return k
