for _ in range(int(input())):
  a, b = input().split()
  if a == b:
    print('=')
  else:
    if 'L' in a and 'L' in b:
      if a.count('X') > b.count('X'):
        print('>')
      else:
        print('<')
    elif 'S' in a and 'S' in b:
      if a.count('X') > b.count('X'):
        print('<')
      else:
        print('>')
    else:
      if a == 'M' or b == 'M':
        if 'S' in a or 'S' in b:
          if 'S' in a:
            print('<')
          else:
            print('>')
        else:
          if 'L' in a:
            print('>')
          else:
            print('<')
      else:
        if 'L' in a and 'S' in b:
          print('>')
        else:
          print('<')
