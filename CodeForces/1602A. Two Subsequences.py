for _ in range(int(input())):
  string = input()
  _min = 999999
  for i in range(len(string)):
      if ord(string[i]) < _min:
          _min = ord(string[i])
          letter = string[i]
          index = i
      if _min <= 97:
          break
  print(f"{string[index]} {string[:index] + string[index+1:]}")
