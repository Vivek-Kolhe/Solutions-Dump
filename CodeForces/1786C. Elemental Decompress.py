import heapq

for _ in range(int(input())):
    n = int(input())
    arr = list(map(int, input().split()))
    cnt = {}
    flag = False
    for i in range(1, n + 1):
        cnt[i] = 0
    for ele in arr:
        cnt[ele] += 1
        if cnt[ele] > 2:
            flag = True
            print("NO")
            break
    if flag:
        continue
    p, q = [0] * n, [0] * n
    pq = []
    v, pairs = [False] * (n + 1), [0] * (n + 1)
    for i in range(n, 0, -1):
        if cnt[i] == 0:
            heapq.heappush(pq, -i)
    for i in range(n, 0, -1):
        if cnt[i] == 1:
            pairs[i] = i
        elif cnt[i] == 2:
            if -pq[0] < i:
                pairs[i] = -pq[0]
                pairs[-pq[0]] = i
                heapq.heappop(pq)
    for i in range(n):
        if not v[arr[i]]:
            p[i] = arr[i]
            q[i] = pairs[arr[i]]
            v[arr[i]] = True
        else:
            p[i] = pairs[arr[i]]
            q[i] = arr[i]
    for i in range(n):
        if p[i] == 0 or q[i] == 0:
            print("NO")
            break
    else:
        print("YES")
        print(*p)
        print(*q)
