for _ in range(int(input())):
  n = int(input())
  arr = list(map(int, input().split()))
  res = "1"
  new = [arr[0]]
  flag = False
  for i in range(1, n):
    if new[-1] <= arr[i] and not flag:
      new.append(arr[i])
      res += "1"
      continue
    if new[0] >= arr[i] and not flag:
      new.append(arr[i])
      res += "1"
      flag = True
      continue
    if new[-1] <= arr[i] and new[0] >= arr[i] and flag:
      new.append(arr[i])
      res += "1"
      continue
    res += "0"
  print(res)
