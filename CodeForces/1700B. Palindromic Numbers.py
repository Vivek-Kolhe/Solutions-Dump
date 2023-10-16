for _ in range(int(input())):
  n = int(input())
  num = int(input())
  checker = (10 ** (n - 1)) * 9
  topal = "9" * n
  if num >= checker:
    topal = "1" * (n + 1)
  print(int(topal) - num)
