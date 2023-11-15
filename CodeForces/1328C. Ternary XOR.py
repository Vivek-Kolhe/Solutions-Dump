for _ in range(int(input())):
  n = int(input())
  x = input()
  a, b = "", ""
  for i in range(n):
    if x[i] == "1":
      a += "1"
      b += "0"
      break
    else:
      half = int(x[i]) // 2
      a += str(half)
      b += str(half)
  a += (n - len(a)) * "0"
  b += x[len(b):]
  print(a)
  print(b)
