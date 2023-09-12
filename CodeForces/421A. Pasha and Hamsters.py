n, ar, al = map(int, input().split())
arthur = list(map(int, input().split()))
alex = list(map(int, input().split()))
res = [1] * n
for ele in alex:
  res[ele-1] = 2
print(*res)
