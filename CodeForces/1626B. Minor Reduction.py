for _ in range(int(input())):
  n = input()
  l = len(n)
  if l == 2:
    print(int(n[0]) + int(n[1]))
  else:
    for i in range(l - 1, 0, -1):
      s = int(n[i]) + int(n[i-1])
      if s >= 10:
        new = n[:i-1] + str(s)
        if i - 1 != l - 2:
          new += n[i+1:]
        print(new)
        break
    else:
      print(str(int(n[0]) + int(n[1])) + n[2:])
