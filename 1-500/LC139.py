def wordBreak(s, wordDict) -> bool:
    d = [False] * (len(s) + 1)
    d[0] = True
    wordDict = set(wordDict)
    for i in range(len(s)):
        if not d[i]:
            continue
        for j in range(1, 20):
            if i + j > len(s):
                break
            if s[i:i + j] in wordDict:
                d[i + j] = True
    return d[-1]
