for _ in range(int(input())):
    n, k = map(int, input().split())
    arr = list(map(int, input().split()))
    flag = False
    arr_ = list(set(arr))
    for i in range(len(arr_)):
        temp = k - arr_[i]
        if temp != arr_[i]:
            if temp in arr_:
                flag = True
                print("Yes")
                break
    if not flag:
        print("No")
