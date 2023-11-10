letters = "abcd"
n = int(input())
for i in range(n):
  print(letters[i%4], end="")
print()
