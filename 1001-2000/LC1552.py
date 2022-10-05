class Solution:
    def maxDistance(self, position: List[int], m :int) -> int:
        def check(x: int) -> bool:
            pre = position[0]
            cnt =1 
            for i in range(1, len(position)):
                if position[i] - pre >= x:
                    pre = position[i]
                    cnt+=1
            return cnt>=m
        
        position.sort()
        left, right, ans = 1, position[-1]-position[0],-1
        while left <= right:
            mid = (left+right)>>1
            if check(mid):
                ans = mid
                left = mid+1
            else:
                right = mid -1
        return ans