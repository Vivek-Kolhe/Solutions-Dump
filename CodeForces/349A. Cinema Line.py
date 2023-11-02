n = int(input())
arr = list(map(int, input().split()))
h = {
    25 : 0,
    50 : 0,
    100 : 0
  }
for i in range(n):
  h[arr[i]] += 1
  change = arr[i] - 25
  if change == 25:
    if h[25] > 0:
      h[25] -= 1
      continue
    else:
      print("NO")
      break
  elif change == 75:
    if h[25] > 0 and h[50] > 0:
      h[25] -= 1
      h[50] -= 1
      continue
    elif h[25] > 2:
      h[25] -= 3
      continue
    else:
      print("NO")
      break
else:
  print("YES")
