n = int(input())
arr = list(set(sorted(list(map(int, input().split())))))
flag = False
for i in range(len(arr) - 2):
  if arr[i + 1] - arr[i] == 1 and arr[i + 2] - arr[i + 1] == 1:
    flag = True
    print("YES")
    break
if not flag:
  print("NO")
