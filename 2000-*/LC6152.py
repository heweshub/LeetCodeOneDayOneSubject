class Solution:
    def minNumberOfHours(self, eng: int, exp: int, energy: List[int], experience: List[int]) -> int:
        ans = 0
        for x, y in zip(energy, experience):
            if eng <= x:
                ans += x + 1 - eng
                eng = x + 1  # 补到刚好超过
            eng -= x
            if exp <= y:
                ans += y + 1 - exp
                exp = y + 1  # 补到刚好超过
            exp += y
        return ans

