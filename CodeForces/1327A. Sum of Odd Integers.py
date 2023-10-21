for _ in range(int(input())):
  n, k = map(int, input().split())
  if k ** 2 > n:
    print("NO")
  elif k % 2 != n % 2:
    print("NO")
  else:
    print("YES")
