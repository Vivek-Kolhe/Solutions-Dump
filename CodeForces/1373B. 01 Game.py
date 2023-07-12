for _ in range(int(input())):
  string = input()
  count = 0
  while True:
    if '01' in string:
      string = string.replace('01', '', 1)
    elif '10' in string:
      string = string.replace('10', '', 1)
    else:
      break
    count += 1
  if count % 2 == 0:
    print('NET')
  else:
    print('DA')
