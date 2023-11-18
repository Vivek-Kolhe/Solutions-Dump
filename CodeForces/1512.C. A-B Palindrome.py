import sys

def solve():
   a, b = map(int, input().split())
   s = input()
 
   if a + b != len(s):
      return -1
   
   res = ['?'] * len(s)
   i, j = 0, len(s) - 1
   q_num = 0
   while i < j:
      x = s[i] + s[j]
      if x in ["01", "10"]:
         return -1
      elif x in ["?0", "0?", "00"]:
         res[i] = res[j] = "0"
         a -= 2
      elif x in ["?1", "1?", "11"]:
         res[i] = res[j] = "1"
         b -= 2
      else: # x == "??"
         q_num += 2
 
      i += 1
      j -= 1
   
   if i == j:
      if s[i] == "0":
         a -= 1
      elif s[i] == "1":
         b -= 1
      else:
         q_num += 1
      res[i] = s[i]
 
   if a < 0 or b < 0:
      return -1
   
   i, j = 0, len(s) - 1
   while i < j:
      if res[i] == res[j] == "?":
         if a >= 2:
            res[i] = res[j] = "0"
            a -= 2
         elif b >= 2:
            res[i] = res[j] = "1"
            b -= 2
         else:
            return -1
      i += 1
      j -= 1
 
   if i == j and res[i] == "?":
      if a >= 1:
         res[i] = "0"
      elif b >= 1:
         res[i] = "1"
      else:
         return -1
 
   return "".join(res)
 
t = int(input())
for _ in range(t):
   print(solve())
