for _ in range(int(input())):
  n = int(input())
  pw = 10 ** (len(str(n)) - 1)
  print(n - pw)
