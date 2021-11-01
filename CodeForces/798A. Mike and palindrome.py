string = input()
unique = list(set(string))
if string == string[::-1] and len(string) % 2 == 0:
  print("NO")
elif string == string[::-1] and len(string) % 2 != 0:
  print("YES")
else:
  flag = False
  for i in range(len(unique)):
    for j in range(len(string)):
      temp = string.replace(string[j], unique[i], 1)
      if temp == temp[::-1]:
        flag = True
        print("YES")
        break
    if flag:
      break
  if not flag:
    print("NO")
