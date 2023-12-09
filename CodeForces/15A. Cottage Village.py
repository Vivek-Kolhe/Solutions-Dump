n, d = map(int, input().split())
arr = []
for _ in range(n):
  arr.append(list(map(int, input().split())))
arr.sort()
cnt = 2
for i in range(n - 1):
  dist = (arr[i+1][0] - (arr[i+1][1] / 2)) - (arr[i][0] + (arr[i][1] / 2))
  if dist > d:
    cnt += 2
    continue
  if dist == d:
    cnt += 1
print(cnt)
