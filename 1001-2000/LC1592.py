def reorderSpaces(text: str) -> str:
    cnt = 0
    for i in text:
        if i == " ":
            cnt += 1
    ans = [i for i in text.split(' ') if i != '']
    ss = cnt // (len(ans) - 1)
    pp = cnt % (len(ans) - 1)
    ans = (ss * " ").join(ans)
    ans += (pp * " ")
    # print(ans)
    return ans


if __name__ == '__main__':
    reorderSpaces("  this   is  a sentence ")
