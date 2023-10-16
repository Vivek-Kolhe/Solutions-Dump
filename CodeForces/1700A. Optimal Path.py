for _ in range(int(input())):
  n, m = map(int, input().split())
  row = (m * (1 + m)) // 2
  col = (n * (m + (n * m))) // 2
  print(row + col - m)
