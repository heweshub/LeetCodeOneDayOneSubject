def garbageCollection(garbage: List[str],travel: List[int]) -> int:
    ans = 0
    right = {}
    for i, s in enumerate(garbage):
        ans += len(s)
        for c in s:
            right[c] = i
    print(right)
    return ans + sum(sum(travel[:r]) for r in right.values())

if __name__ == "__main__":
    garbageCollection(["G","P","GP","GG"], [2,4,3])