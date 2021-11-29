for _ in range(int(input())):
  string = input()
  if "a" + string != ("a" + string)[::-1]:
    print("YES")
    print("a" + string)
  elif string + "a" != (string + "a")[::-1]:
    print("YES")
    print(string + "a")
  else:
    print("NO")
