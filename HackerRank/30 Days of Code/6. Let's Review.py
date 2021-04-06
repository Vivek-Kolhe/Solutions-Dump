T = int(input())

for tc in range(T):
    string = input()
    even = ""
    odd = ""
    for i in range(len(string)):
        if i % 2 == 0:
            even += string[i]
        else:
            odd += string[i]
    print(even, odd)
