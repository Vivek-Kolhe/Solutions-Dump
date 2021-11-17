for _ in range(int(input())):
  n = int(input())
  string = input()
  temp, flag = list(set(string)), False
  for i in range(1, len(temp)):
    if string.rfind(temp[i-1]) > string.find(temp[i]):
      flag = True
      print("NO")
      break
  if not flag:
    print("YES")
