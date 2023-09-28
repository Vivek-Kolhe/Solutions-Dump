# cook your dish here
for _ in range(int(input())):
    cost = int(input())
    com = cost * 0.2
    if 100 % com == 0:
        print(int(100 // com))
    else:
        print(int(100 // com) + 1)
