for _ in range(int(input())):
  s = input()
  n = len(s)
  curr = 0
  cnt = 0
  if n == 1:
    print(1)
  else:
    while curr < n - 1:
      st = []
      for i in range(curr, n):
        if s[i] not in st and len(st) < 3:
          st.append(s[i])
          continue
        if s[i] not in st and len(st) == 3:
          break
      curr = i
      cnt += 1
      if curr == n - 1:
        if s[curr] not in st:
          cnt += 1
    print(cnt)
