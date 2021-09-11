n = int(input())
s = input().upper()
sub_string = ""
res = 0
 
for i in range(n - 1):
  current = 0
  for j in range(n - 1):
    if s[j] == s[i] and s[j + 1] == s[i + 1]:
      current += 1
  if res < current:
    res = current
    sub_string = s[i] + s[i + 1]
print(sub_string)
