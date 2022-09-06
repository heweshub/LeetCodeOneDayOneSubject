from linecache import cache


class Solution:
    def numberOfWays(self, s: int, e: int, k: int) -> int:
        # d = abs(s-e)
        # if (d+k) %2 == 1 or d > k:
        #     return 0
        # x = (d+k) // 2
        # y = (k-d) // 2
        # p = 1e9+7
        # def C(n,m):
        #     a = 1
        #     for i in range(m):
        #         a = a*(n-i)*pow(i+1,p-2,p)%p
        #     return a
        # return C(x+y, y)
        pass
        MOD = 10**9+7
        @cache
        def f(x: int, left: int) ->int:
            if abs(x-e) > left: return 0
            if left == 0: return 1
            return (f(x-1,left-1) + f(x+1,left-1))%MOD
        return f(s,k)