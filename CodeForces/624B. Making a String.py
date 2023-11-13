n = int(input())
arr = list(map(int, input().split()))
arr.sort(reverse=True)
s = arr[0]
prev = arr[0]
cnt = 1
for i in range(1, n):
  if arr[i] == prev:
    arr[i] = max(arr[i] - cnt, 0)
    cnt += 1
  else:
    prev = arr[i]
    cnt = 1
    if prev >= arr[i-1]:
      arr[i] = max(0, arr[i-1] - 1)
    prev = arr[i]
  s += arr[i]
print(s)
