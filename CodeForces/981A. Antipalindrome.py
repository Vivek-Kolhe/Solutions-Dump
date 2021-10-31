string = input()
result = []
for i in range(len(string)):
  temp = string[i:]
  if temp != temp[::-1]:
    result.append(len(temp))
for i in range(len(string)):
  temp = string[:len(string)-1]
  if temp != temp[::-1]:
    result.append(len(temp))
try:
  print(max(result))
except:
  print(0)
