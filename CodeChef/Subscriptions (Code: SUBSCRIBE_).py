# cook your dish here
for _ in range(int(input())):
    n, x = map(int, input().split())
    num = n // 6 if n % 6 == 0 else (n // 6) + 1
    print(num * x)
