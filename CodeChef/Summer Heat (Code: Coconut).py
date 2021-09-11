# Problem Summer Heat, Problem Code: COCONUT

n = int(input())
for _ in range(n):
    xa, xb, Xa, Xb = map(int, input().split())
    print(round(Xa/xa) + round(Xb/xb))
