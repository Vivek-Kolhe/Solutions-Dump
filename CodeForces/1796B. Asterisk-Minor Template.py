for _ in range(int(input())):
  s = input()
  t = input()
  if s == t:
    print("YES")
    print(s)
  elif s[-1] == t[-1]:
    print("YES")
    print("*" + s[-1])
  elif s[0] == t[0]:
    print("YES")
    print(s[0] + "*")
  else:
    sl, tl = len(s), len(t)
    if sl > tl:
      s, t = t, s
      sl, tl = tl, sl
    ans = ""
    for i in range(sl - 1):
      if s[i] + s[i+1] in t:
        print("YES")
        print("*" + s[i] + s[i+1] + "*")
        break
    else:
      print("NO")
