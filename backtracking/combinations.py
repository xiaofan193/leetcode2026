class Solution:
    def combine(self, n: int, k: int) -> List[List[int]]:
        result = []  # Store results set
        self.backtracking(n, k, 1, [], result)
        return result
    def backtracking(self, n, k, startIndex, path, result):
        if len(path) == k:
            result.append(path[:])
            return
        for i in range(startIndex, n - (k - len(path)) + 2):  # Optimized section
            path.append(i)  # Handle node
            self.backtracking(n, k, i + 1, path, result)
            path.pop()  # Backtrack and undo node handling