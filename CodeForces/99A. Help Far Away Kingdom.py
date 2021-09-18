s = input().split(".")
if s[0][-1] != "9" and int(s[-1][0]) < 5:
  print(int(s[0]))
elif s[0][-1] != "9" and int(s[-1][0]) >= 5:
  print(int(s[0]) + 1)
else:
  print("GOTO Vasilisa.")
