for _ in range(int(input())):
  s = input()
  n = len(s)
  curr = 0
  cnt = 0
  while curr < n:
    st = []
    for i in range(curr, n):
      if s[i] not in st and len(st) < 4:
        st.append(s[i])
      if s[i] not in st and len(st) == 3:
        break
    curr = i
    cnt += 1
  print(count)
