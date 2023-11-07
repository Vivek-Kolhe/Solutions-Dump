n, l = map(int, input().split())
arr = list(map(int, input().split()))
arr.sort()
maxi = 0
for i in range(n - 1):
  if arr[i+1] - arr[i] > maxi:
    maxi = arr[i+1] - arr[i]
res = max([maxi / 2, arr[0] - 0, l - arr[n-1]])
print(float(res))
