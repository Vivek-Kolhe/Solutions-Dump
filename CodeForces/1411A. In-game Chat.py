for _ in range(int(input())):
  n = int(input())
  string = input()
  if n == 1:
    if string != ")":
      print("No")
    else:
      print("Yes")
  else:
    count = 0
    for i in range(n - 1):
      if string[i] == ")":
        count += 1
      if string[i + 1] != ")":
        count = 0
    count += 1
    if count > n - count:
      print("Yes")
    else:
      print("No")
