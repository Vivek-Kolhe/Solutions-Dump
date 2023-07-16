s1, s2 = input(), input() 
minimum = min(len(s1), len(s2)) 
i = 1 
while i <= minimum and s1[-i] == s2[-i]:
  i += 1
i -= 1
print(len(s1) + len(s2) - 2 * i)
