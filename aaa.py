# -*- coding: UTF-8 -*-
"""
author: cjhcw
"""


def intersection(nums):
    n = len(nums)
    ans = []
    if n == 1:
        nums[0].sort()
        return nums[0]
    for i in range(len(nums[0])):
        for j in range(1, n):
            if nums[0][i] not in nums[j]:
                break
            if j == n - 1:
                ans.append(nums[0][i])
    print(ans)
    return ans.sort()


def appealSum(s):
    cnt = 0
    n = len(s)
    for i in range(1, n + 1):
        for j in range(n):
            if j + i > n:
                break
            else:
                cnt += dupli(s[j:i + j])
    return cnt


def dupli(ss):
    liS = [ss[i] for i in range(len(ss))]
    return len(set(liS))


if __name__ == '__main__':
    # intersection([[44, 21, 48]])
    # print(dupli("aabd"))
    print(appealSum("abbca"))
    print(appealSum("code"))
