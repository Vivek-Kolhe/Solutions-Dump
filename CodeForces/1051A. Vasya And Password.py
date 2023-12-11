for _ in range(int(input())):
  string = input()
  c, s, d = 0, 0, 0
  cp, sp, dp = -1, -1, -1
  for i in range(len(string)):
    if string[i].isdigit():
      d += 1
      dp = i
    if string[i].isupper():
      c += 1
      cp = i
    if string[i].islower():
      s += 1
      sp = i
  if c > 0 and d > 0 and s > 0:
    print(string)
  elif (sp == -1 and dp == -1) or (dp == -1 and cp == -1) or (cp == -1 and sp == -1):
    if s > 0 and (c == 0 and d == 0):
      print(string[0].upper() + "6" + string[2:])
    if c > 0 and (d == 0 and s == 0):
      print(string[0].lower() + "6" + string[2:])
    if d > 0 and (s == 0 and c == 0):
      print("aB" + string[2:])
  else:
    string = list(string)
    if d == 0 and (s > 0 and c > 0):
      if s > c:
        string[sp] = "5"
      else:
        string[cp] = "5"
    if s == 0 and (c > 0 and d > 0):
      if c > d:
        string[cp] = string[cp].lower()
      else:
        string[dp] = "a"
    if c == 0 and (s > 0 and d > 0):
      if s > d:
        string[sp] = string[sp].upper()
      else:
        string[dp] = "A"
    print("".join(string))
