# cook your dish here
for _ in range(int(input())):
    n, m = map(int, input().split())
    x, y = map(int, input().split())
    r = (n - x) * m
    l = (x - 1) * m
    u = (y - 1) * n
    d = (m - y) * n
    print(max([r, l, u, d]))
