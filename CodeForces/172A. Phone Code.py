arr = []
for _ in range(int(input())):
  arr.append(input())
first = arr[0]
flag = True
res = ""
for i in range(1, len(first) + 1):
  for j in range(1, len(arr)):
    if not arr[j].startswith(first[:i]):
      flag = False
      break
  if not flag:
    res = first[:i-1]
    break
print(len(res))
