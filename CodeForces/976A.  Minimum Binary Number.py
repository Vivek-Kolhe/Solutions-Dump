n = int(input())
string = input()
string = "".join(sorted(string))[::-1]
for i in range(string.count("1") - 1):
  string = string.replace("11", "1")
print(string)
