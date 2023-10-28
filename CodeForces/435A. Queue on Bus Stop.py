n, m = map(int, input().split())
arr = list(map(int, input().split()))
cnt = 0
i = 0
while i < n:
  c = m
  for j in range(i, n):
    if c >= arr[j]:
      c -= arr[j]
    else:
      break
    i += 1
  cnt += 1
print(cnt)
