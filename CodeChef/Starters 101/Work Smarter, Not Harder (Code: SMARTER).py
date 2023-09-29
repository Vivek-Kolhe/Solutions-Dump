# cook your dish here
def round_(n, m):
    if n % m == 0:
        return n // m
    return (n // m) + 1

for _ in range(int(input())):
    l, t, h = map(int, input().split())
    tt = round_(l, t)
    ht = round_(l, h)
    if tt == ht:
        print(-1)
    else:
        print(tt - ht - 1)
