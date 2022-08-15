for _ in range(int(input())):
  string = input()
  if len(string) % 2 != 0:
    print("NO")
  else:
    mid = len(string) // 2
    if string[:mid] == string[mid:]:
      print("YES")
    else:
      print("NO")
