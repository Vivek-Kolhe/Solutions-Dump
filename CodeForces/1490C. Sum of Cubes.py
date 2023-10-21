for _ in range(int(input())):
  n = int(input())
  flag = True
  low, high = 1, round(n ** (1/3))
  while low <= high:
    mid = (low ** 3) + (high ** 3)
    if mid == n:
      print("YES")
      flag = False
      break
    elif mid < n:
      low += 1
    else:
      high -= 1
  if flag:
    print("NO")
