for _ in range(int(input())):
  s = input()
  sm, cap, ans = [], [], []
  n = len(s)
  for i in range(n):
    if s[i] != "b" and s[i] != "B":
      ans.append(s[i])
      if s[i].islower():
        sm.append(len(ans) - 1)
      elif s[i].isupper():
        cap.append(len(ans) - 1)
    else:
      if s[i] == "b":
        if sm:
          ind = sm.pop()
          ans[ind] = ""
      elif s[i] == "B":
        if cap:
          ind = cap.pop()
          ans[ind] = ""
  print("".join(ans))
