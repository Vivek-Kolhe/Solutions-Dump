import heapq

n, m = map(int, input().split())
arr = list(map(int, input().split()))
curr = 0
ans, heap = [], []
for i in range(n):
    curr += arr[i]
    if curr <= m:
        heapq.heappush(heap, -arr[i])
        ans.append(0)
    else:
        temp = curr
        pops = []
        c = 0
        while temp > m:
            p = heapq.heappop(heap)
            temp += p
            pops.append(p)
            c += 1
        for ele in pops:
            heapq.heappush(heap, ele)
        heapq.heappush(heap, -arr[i])
        ans.append(c)
print(*ans)
