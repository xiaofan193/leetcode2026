class Solution:
    def plusOne(self, digits: List[int]) -> List[int]:
        n = len(digits)
        ans = []

        k = 0

        if digits[n-1] ==9:
            ans.append(0)
            k =1
        else:
            ans.append(digits[n-1] + 1)
            k =0

        for i in range (n-2,-1,-1):
            if digits[i] + k == 10 :
                ans.append(0)
                k = 1
            else:
                ans.append(digits[i] + k)
                k = 0

        if k == 1:
            ans.append(1)


        ans.reverse()

        return ans
            
        
            