for _ in range(int(input())):
  t = input()
  n = len(t)
  if t.count("1") == n or t.count("0") == n:
    print(t, end="")
  else:
    for i in range(n - 1):
      if t[i] == t[i+1]:
        alt = "1" if t[i] == "0" else "0"
        print(t[i] + alt, end="")
      else:
        print(t[i], end="")
    print(t[-1], end="")
  print()
