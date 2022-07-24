for _ in range(int(input())):
  string_one = input()
  string_two = input()
  flag = True
  ind = [i for i in range(len(string_one)) if string_one[i] == string_two]
  for i in ind:
    if i % 2 == 0:
      print("YES")
      flag = False
      break
  if flag:
    print("NO")
