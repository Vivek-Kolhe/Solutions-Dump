n = int(input())
if n <= 9:
  print(1)
else:
  pow = 10 ** (len(str(n)) - 1)
  next = n + pow - (n % pow)
  print(next - n)
