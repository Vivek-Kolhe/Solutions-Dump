n, k = map(int, input().split())
arr = list(map(int, input().split()))
s = sum(arr)
for i in range(n - 1):
  if abs(arr[i] + arr[i+1]) < k:
    arr[i+1] = k - arr[i]
print(sum(arr) - s)
print(*arr)
