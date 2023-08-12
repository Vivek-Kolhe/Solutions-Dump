for _ in range(int(input())):
  n = int(input())
  end = n - 2
  length = (end * (end + 1)) + (4 * n) + end + 2
  print(length)
