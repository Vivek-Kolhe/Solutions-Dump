vowels = ["A", "E", "I", "O", "U", "Y"]
s = input()
prev = 0
ability = 0
s = "A" + s + "A"
n = len(s)
for i in range(n):
  if s[i] in vowels:
    ability = max(ability, i - prev)
    prev = i
print(ability if prev != 0 else n + 1)
