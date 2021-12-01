import string
alphabets = string.ascii_lowercase
for _ in range(int(input())):
  n, a, b = map(int, input().split())
  s = (alphabets[:b] * ((n // b) + 1))[:n]
  print(s)
