for _ in range(int(input())):
  n = int(input())
  arr = list(map(int, input().split()))
  flag = True
  for i in range(n - 1):
    diff = abs(arr[i] - arr[i+1])
    if not (diff == n - 1 or diff == 1):
      flag = False
      print("NO")
      break
  if flag:
    diff = abs(arr[0] - arr[-1])
    if not (diff == n - 1 or diff == 1):
      print("NO")
    else:
      print("YES")
