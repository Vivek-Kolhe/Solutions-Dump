n = int(input())
s = "bbaac"
for i in range(n):
  print(s[i%4], end="")
print()
