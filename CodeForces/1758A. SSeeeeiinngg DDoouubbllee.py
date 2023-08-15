for _ in range(int(input())):
  s = list(input())
  res = []
  for char in s:
    res.insert(0, char)
    res.insert(len(res), char)
  print("".join(res))
