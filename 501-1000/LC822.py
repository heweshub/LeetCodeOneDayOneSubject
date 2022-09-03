import itertools


class Solution(object):
    def flipgame(self, fronts, backs):
        same = {x for i, x in enumerate(fronts) if x == backs[i]} 
        ans = 9999
        for x in itertools.chain(fronts,backs):
            print(x)
            if x not in same:
                ans = min(ans, x)
        return ans % 9999

