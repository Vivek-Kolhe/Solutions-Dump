for _ in range(int(input())):
  a, b, n, s = map(int, input().split())
  ra, rb = s // n, s % n
  if ra <= a and rb <= b:
    print("YES")
  else:
    ones = s - (a * n)
    if ra > a and ones <= b:
      print("YES")
    else:
      print("NO")
