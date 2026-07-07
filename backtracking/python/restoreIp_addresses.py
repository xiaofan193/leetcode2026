class Solution:
    def restoreIpAddresses(self, s: str) -> List[str]:
        result = []
        self.backtracking(s, 0, 0, "", result)
        return result

    def backtracking(self, s, start_index, point_num, current, result):
        if point_num == 3:  # If the number of dots is 3, segmentation ends
            if self.is_valid(s, start_index, len(s) - 1):  # Check if the fourth segment is valid
                current += s[start_index:]  # Add the last segment
                result.append(current)
            return

        for i in range(start_index, len(s)):
            if self.is_valid(s, start_index, i):  # Check if the [start_index, i] segment is valid
                sub = s[start_index:i + 1]
                self.backtracking(s, i + 1, point_num + 1, current + sub + '.', result)
            else:
                break

    def is_valid(self, s, start, end):
        if start > end:
            return False
        if s[start] == '0' and start != end:  # Segments starting with 0 are invalid
            return False
        num = 0
        for i in range(start, end + 1):
            if not s[i].isdigit():  # Non-digit characters are invalid
                return False
            num = num * 10 + int(s[i])
            if num > 255:  # Numbers greater than 255 are invalid
                return False
        return True