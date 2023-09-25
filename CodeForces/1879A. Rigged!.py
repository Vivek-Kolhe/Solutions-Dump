for _ in range(int(input())):
  n = int(input())
  arr = []
  for i in range(n):
    arr.append(list(map(int, input().split())))
  for i in range(1, n):
    if arr[i][0] >= arr[0][0] and arr[i][-1] >= arr[0][-1]:
      print(-1)
      break
  else:
    print(arr[0][0])
