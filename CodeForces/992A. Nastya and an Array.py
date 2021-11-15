n = int(input())
arr = list(set(list(map(int, input().split()))))
if 0 in arr:
  print(len(arr) - 1)
else:
  print(len(arr))
