for _ in range(int(input())):
  n = int(input())
  s = input()
  for i in range(n // 2):
    if s[i] == s[-i-1]:
      continue
    else:
      l, r = s[i], s[-1-i]
      if ord(l) - 1 == ord(r) + 1 or ord(l) + 1 == ord(r) - 1:
        continue
      else:
        print("NO")
        break
  else:
    print("YES")
