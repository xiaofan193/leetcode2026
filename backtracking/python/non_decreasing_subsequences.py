class Solution:
    def findSubsequences(self, nums):
        result = []
        path = []
        self.backtracking(nums, 0, path, result)
        return result
    
    def backtracking(self, nums, startIndex, path, result):
        if len(path) > 1:
            result.append(path[:])  # Use slices for a copy of the current path to append to result set
            # Note don't add return here, as we need to collect all nodes on the tree
        
        uset = set()  # Use a set to de-duplicate elements at this level
        for i in range(startIndex, len(nums)):
            if (path and nums[i] < path[-1]) or nums[i] in uset:
                continue
            
            uset.add(nums[i])  # Record this element as used at this level
            path.append(nums[i])
            self.backtracking(nums, i + 1, path, result)
            path.pop()