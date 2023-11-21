for _ in range(int(input())):
  n = int(input())
  l = n + 2
  arr = list(map(int, input().split()))
  arr.sort()
  s = 0
  for i in range(n):
    s += arr[i]
  if s == arr[n] or s == arr[n+1]:
    for i in range(n):
      print(arr[i], end=" ")
    print()
    continue
  s += arr[n]
  ind = -1
  flag = False
  for i in range(n + 1):
    if arr[n+1] == s - arr[i]:
      ind = i
      flag = True
      break
  if not flag:
    print(-1)
  else:
    for i in range(n + 1):
      if i == ind:
        continue
      else:
        print(arr[i], end=" ")
    print()
