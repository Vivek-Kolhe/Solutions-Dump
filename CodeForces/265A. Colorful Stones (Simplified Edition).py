s = input()
ins = input()
curr = 0
for i in range(len(ins)):
  if ins[i] == s[curr]:
    curr += 1
print(curr + 1)
