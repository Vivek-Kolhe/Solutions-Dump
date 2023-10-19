n = int(input())
arr = []
for _ in range(n):
  arr.append(list(map(int, input().split())))
if n < 2:
  print("Poor Alex")
else:
  arr.sort()
  for i in range(n - 1):
    curr, next = arr[i], arr[i+1]
    if next[0] > curr[0] and next[-1] < curr[-1]:
      print("Happy Alex")
      break
  else:
    print("Poor Alex")
