# cook your dish here
mod = (10 ** 9) + 7
for _ in range(int(input())):
    n = int(input())
    num = input()
    count = 1
    for i in range(0, n - 1, 2):
        res = num[i+2]
        a, b = num[i], num[i+1]
        if res == "1":
            if a == "0" and b == "0":
                count = 0
            else:
                count = count * 2 % mod
        else:
            if a == "0" and b == "0":
                count = count * 3 % mod
    print(count)
