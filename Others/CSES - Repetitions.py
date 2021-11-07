string = input()
count, temp = 1, 1
for i in range(len(string) - 1):
    if string[i] == string[i + 1]:
        temp += 1
    else:
        temp = 1
    if temp > count:
        count = temp
print(count)