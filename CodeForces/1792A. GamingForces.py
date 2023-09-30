for _ in range(int(input())):
  n = int(input())
  arr = list(map(int, input().split()))
  arr.sort()
  count = 0
  for i in range(n):
    if arr[i] == 0:
      continue
    elif arr[i] == 1:
      count += 1
      if i + 1 < n:
        arr[i+1] -= 1
    else:
      arr[i] = 0
      count += 1
  print(count)
