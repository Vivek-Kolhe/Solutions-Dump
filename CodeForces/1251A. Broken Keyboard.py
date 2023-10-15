for _ in range(int(input())):
  s = input()
  n = len(s)
  h = {}
  i = 0
  while i < n:
    j = i
    while j < n and s[j] == s[i]:
      j += 1
    j -= 1
    if (j - i) % 2 == 1:
      if h.get(s[i]) is None:
        h[s[i]] = False
    else:
      h[s[i]] = True
    i = j + 1
  keys = sorted(h.keys())
  res = ""
  for key in keys:
    if h[key]:
      res += key
  print(res)
