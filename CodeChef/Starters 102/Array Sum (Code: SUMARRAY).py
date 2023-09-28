for _ in range(int(input())):
    n, k = map(int, input().split())
    if k < (n * 2) - (n // 2): print(-1)
    else:
        min_arr = [1, 2] * (n // 2)
        rem = k - sum(min_arr)
        if rem % 2 == 1: print(-1)
        else:
            rounds = rem // (n * 2)
            for i in range(n):
                min_arr[i] += rounds * 2
                rem -= rounds * 2
            for i in range(n):
                if rem == 0:
                    break
                min_arr[i] += 2
                rem -= 2
            for i in range(n):
                if min_arr[i] > 10 ** 5:
                    print(-1)
                    break
            else:
                print(*min_arr)
