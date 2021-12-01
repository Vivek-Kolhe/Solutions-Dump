keyboard = "qwertyuiopasdfghjkl;zxcvbnm,./"
direction = input()
string = input()
ans = ""
flag = 1 if direction == "R" else 0
for i in range(len(string)):
  ind = keyboard.find(string[i]) - 1 if flag else keyboard.find(string[i]) + 1
  ans += keyboard[ind]
print(ans)
