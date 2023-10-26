for _ in range(int(input())):
  p = "".join(sorted(input()))
  h = input()
  pn, hn = len(p), len(h)
  for i in range(hn - pn + 1):
    temp = h[i:i+pn]
    if p == "".join(sorted(temp)):
      print("YES")
      break
  else:
    print("NO")
