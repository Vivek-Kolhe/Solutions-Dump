string = input()
flag_l, flag_s, flag_d = False, False, False
if len(string) > 4:
  for i in range(len(string)):
    if string[i].isupper():
      flag_l = True
    elif string[i].islower():
      flag_s = True
    elif string[i].isnumeric():
      flag_d = True
    if flag_l and flag_s and flag_d:
      print("Correct")
      break
  if not flag_l or not flag_s or not flag_d:
    print("Too weak")
else:
  print("Too weak")
