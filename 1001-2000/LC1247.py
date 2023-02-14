class Solution:
    def minimumSwap(self, s1: str, s2: str) -> int:
        cn1, cn2 = 0, 0
        for i in range(len(s1)):
            if s1[i] == 'x' and s2[i] == 'y':
                cn1 += 1
            elif s1[i] == 'y' and s2[i] == 'x':
                cn2 += 1
        if (cn1 + cn2) % 2 != 0:
            return -1
        n1, m1 = divmod(cn1, 2)
        n2, m2 = divmod(cn2, 2)
        return n1 + n2 + 2 * m1
