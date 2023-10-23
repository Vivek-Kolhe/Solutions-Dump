n = int(input())
arr = list(map(int, input().split()))
arr.sort()
if arr[0] == 1:
  print(-1)
else:
  if arr[0] > 1:
    print(1)
