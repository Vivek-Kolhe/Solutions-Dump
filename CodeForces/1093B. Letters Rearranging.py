for _ in range(int(input())):
  string = input()
  if string.count(string[0]) == len(string):
    print(-1)
  else:
    print("".join(sorted(string)))
