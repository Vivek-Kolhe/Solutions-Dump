n = int(input())
arr = list(map(int, input().split()))
ans = (n - 1) // 2
res = [0] * n
arr.sort()
ind = 0
for i in range(1, n, 2):
  res[i] = arr[ind]
  ind += 1
for i in range(0, n, 2):
  res[i] = arr[ind]
  ind += 1
print(ans)
print(*res)
