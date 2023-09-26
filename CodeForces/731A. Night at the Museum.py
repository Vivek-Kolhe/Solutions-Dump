import string
alphabets = string.ascii_lowercase

def comb(n, target):
  if n > target:
    return min(abs(n - target), abs(26 - n + target))
  return min(abs(n - target), abs(26 + n - target))

s = input()
count = 0
curr = "a"
for i in range(len(s)):
  count += comb(alphabets.find(curr), alphabets.find(s[i]))
  curr = s[i]
print(count)
