import string
alphabets = string.ascii_lowercase
n, k = map(int, input().split())
temp = alphabets[:k]
while True:
  if len(temp) >= n:
    break
  temp += temp
print(temp[:n])
