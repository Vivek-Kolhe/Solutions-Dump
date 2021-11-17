for _ in range(int(input())):
  n = int(input())
  arr = sorted(list(map(int, input().split())))
  flag = True
  for i in range(n - 1):
    if not abs(arr[i] - arr[i + 1]) <= 1:
      flag = False
      print("NO")
      break
  if flag:
    print("YES")
