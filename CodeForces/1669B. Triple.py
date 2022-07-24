for _ in range(int(input())):
  n = int(input())
  arr = list(map(int, input().split()))
  arr.sort()
  flag = False
  for i in range(n - 2):
    if arr[i] == arr[i + 1] and arr[i + 1] == arr[i + 2]:
      flag = True
      print(arr[i])
      break
  if not flag:
    print(-1)
