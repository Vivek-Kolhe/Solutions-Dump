n, m = map(int, input().split())
arr = [list(map(int, input().split())) for _ in range(m)]
arr.sort(key=lambda x: x[-1], reverse=True)
count = 0
for i in range(m):
  if n >= arr[i][0]:
    count += (arr[i][0] * arr[i][-1])
    n -= arr[i][0]
  else:
    temp = arr[i][0] - n
    count += ((arr[i][0] - temp) * arr[i][-1])
    n -= arr[i][0] - temp
  if n == 0:
    break
print(count)
