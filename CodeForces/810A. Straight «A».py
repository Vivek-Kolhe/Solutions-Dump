n, k = map(int, input().split())
arr = list(map(int, input().split()))
need = k - 0.5
s = sum(arr)
avg = s / n
cnt = 0
while avg < need:
  cnt += 1
  s += k
  n += 1
  avg = s / n
print(cnt)
