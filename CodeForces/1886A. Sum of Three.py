for _ in range(int(input())):
  n = int(input())
  if n < 7 or n == 9:
    print("NO")
  else:
    res = [1]
    if n % 3 == 1 or n % 3 == 2:
      res.append(2)
      res.append(n - 3)
    else:
      res.append(4)
      res.append(n - 5)
    print("YES")
    print(*res)
