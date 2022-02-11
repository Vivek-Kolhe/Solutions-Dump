for _ in range(int(input())):
  n, k = map(int, input().split())
  string = input()1634A. Reverse and Concatenate.py
  if k == 0 or string == string[::-1]:
    print(1)
  else:
    print(2)
