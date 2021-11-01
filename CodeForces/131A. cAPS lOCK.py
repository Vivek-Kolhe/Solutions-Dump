string = input()
if string[0] == string[0].lower() and string[1:] == string[1:].upper():
  print(string.capitalize())
elif string == string.upper():
  print(string.lower())
else:
  print(string)
