# cook your dish here
for _ in range(int(input())):
    x, y = map(int, input().split())
    res = x // y
    print(20 if res > 20 else res)
