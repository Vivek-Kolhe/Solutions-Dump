for _ in range(int(input())):
    n = int(input())
    if n > 1099 or n % 11 == 0 or n % 111 == 0:
        print("YES")
        continue
    flag = True
    while n > 0:
        if n % 11 == 0:
            flag = False
            print("YES")
            break
        n -= 111
    if flag:
        if n == 0:
            print("YES")
        else:
            print("NO")
