string = input()
for i in range(5, 0, -1):
  string = string.replace('C' * i, 'x')
  string = string.replace('P' * i, 'y')
print(len(string))
