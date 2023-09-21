# cook your dish here
for _ in range(int(input())):
    l, r, m = map(int, input().split())
    pages = m // l if m % l == 0 else (m // l) + 1
    code = m // r
    print(code + pages)
