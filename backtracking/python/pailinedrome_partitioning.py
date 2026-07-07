class Solution:
    def partition(self, s: str) -> List[List[str]]:
        '''
        Recursion helps to traverse vertically
        For-loop helps to traverse horizontally
        When the partition line iterates to the end of the string, it means one way is found
        Similar to combination problems, to avoid duplicate partitioning at the same position, `start_index` is needed to mark the starting position for the next round of recursion (partition line)
        '''
        result = []
        self.backtracking(s, 0, [], result)
        return result

    def backtracking(self, s, start_index, path, result ):
        # Base Case
        if start_index == len(s):
            result.append(path[:])
            return
        
        # Single Level Recursion Logic
        for i in range(start_index, len(s)):
            # A step unique to this compared to other combination problems:
            # Determine whether the substring from [start_index, i] is a palindrome
            if self.is_palindrome(s, start_index, i):
                path.append(s[start_index:i+1])
                self.backtracking(s, i+1, path, result)   # Recursion for vertical traversal: partition from the next spot, and determine the rest is still a palindrome
                path.pop()             # Backtracking

    def is_palindrome(self, s: str, start: int, end: int) -> bool:
        i: int = start        
        j: int = end
        while i < j:
            if s[i] != s[j]:
                return False
            i += 1
            j -= 1
        return True