for _ in range(int(input())):
  n = int(input())
  string = input()
  result = ""
  i, j = 0, 0
  while i < n:
    result += string[i]
    for j in range(i+1, n):
      if string[j] == string[i]:
        i = j + 1
        break
  print(result)
