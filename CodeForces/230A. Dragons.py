arr, flag = [], True
s, n = map(int, input().split())
for _ in range(n):
  arr.append(list(map(int, input().split())))
arr.sort()
for i in range(n):
  if s > arr[i][0]:
    s += arr[i][-1]
  else:
    flag = False
    break
if flag:
  print("YES")
else:
  print("NO")
