vowels = "aeiou"
word = input()
i = 0
flag = True
while i < len(word):
  if word[-1] not in vowels and word[-1] != "n":
    flag  = False
    print("NO")
    break
  if word[i] == "n":
    i += 1
  elif word[i] not in vowels:
    if word[i+1] in vowels:
      i += 2
    else:
      flag = False
      print("NO")
      break
  else:
    i += 1
if flag:
  print("YES")
