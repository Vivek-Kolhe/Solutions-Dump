n = int(input())
arr = list(map(int, input().split()))
if 1 in arr:
  hours = arr + arr
  res, count = [], 1
  for i in range((2 * n) - 1):
    if hours[i] == 1 and hours[i+1] == 1:
      count += 1
    else:
      res.append(count)
      count = 1
  print(max(res))
else:
  print(0)
