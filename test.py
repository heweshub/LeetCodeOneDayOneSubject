from secrets import choice


def draw(s1: str,s2: str) -> int:
    n1, n2 = len(s1), len(s2)
    cnt = 0
    i,j = 0,0
    for i ,v in enumerate(s1):
        if j == n2:
            break
        while s2[j] != v and j < n2:
            j+=1
        else:
            cnt+=1
    print(cnt)
    choice(cnt)

if __name__ == '__main__':
    # draw("ATM","CAT")
    a = list(map,)
    
