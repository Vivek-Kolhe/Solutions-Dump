for _ in range(int(input())):
  n = int(input())
  first = input()
  second = input()
  flag = True
  for i in range(n):
    if (first[i] == second[i] == "R") or ((first[i] == "B" or first[i] == "G") and (second[i] == "B" or second[i] == "G")):
      continue
    else:
      flag = False
      print("NO")
      break
  if flag:
    print("YES")
