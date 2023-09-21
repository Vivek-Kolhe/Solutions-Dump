# cook your dish here
from random import shuffle

org = "codechef"
h = {}
for _ in range(int(input())):
    s = sorted(list(input()))
    ss = "".join(s)
    if ss in h:
        print(h[ss])
        continue
    h[ss] = -1
    i = 0
    while i < 300:
        shuffle(s)
        curr = "".join(s)
        flag = True
        for ind in range(8):
            if org[ind] == curr[ind]:
                flag = False
                break
        if flag:
            h[ss] = curr
            break
        i += 1
    print(h[ss])
