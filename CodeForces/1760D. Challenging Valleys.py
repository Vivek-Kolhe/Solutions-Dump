for _ in range(int(input())):
  n = int(input())
  arr = list(map(int, input().split()))
  flag, valley = True, True
  for i in range(1, n):
    if flag and arr[i] > arr[i-1]:
      flag = False
    if not flag and arr[i] < arr[i-1]:
      valley = False
  if valley:
    print("YES")
  else:
    print("NO")
