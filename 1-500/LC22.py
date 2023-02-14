def generateParenthesis(n):
    ans = []

    def backtrack(s, left, right):
        if len(s) == 2 * n:
            ans.append(''.join(s))
            return
        if left < n:
            s.append('(')
            backtrack(s, left + 1, right)
            s.pop()
        if right < left:
            s.append(')')
            backtrack(s, left, right + 1)
            s.pop()

    backtrack([], 0, 0)
    return ans


if __name__ == '__main__':
    print(generateParenthesis(3))
    # s = [1,2,3,4,5,6]
    # s.pop()
    # print(s)
    # s = [1,2,3,4,5,6]
    # s = s[:-1]
    # print(s)
