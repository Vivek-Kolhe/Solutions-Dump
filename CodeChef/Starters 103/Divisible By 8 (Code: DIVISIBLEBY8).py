# cook your dish here
for _ in range(int(input())):
    n, num = map(int, input().split())
    if num % 8 == 0:
        print(num)
    else:
        if num <= 8:
            print(8)
        else:
            for i in range(1, 10):
                res = (num // 10) * 10 + i
                if res % 8 == 0:
                    print(res)
                    break
