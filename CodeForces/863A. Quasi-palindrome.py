num = input()
if num[::-1] == num:
  print("YES")
else:
  num = num[::-1]
  for i in range(len(num)):
    if num[i] != "0":
      new_num = ("0" * i) + num[::-1]
      break
  if new_num != new_num[::-1]:
    print("NO")
  else:
    print("YES")
