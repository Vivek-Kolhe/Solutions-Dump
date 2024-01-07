import heapq

n, k = map(int, input().split())
arr = list(map(int, input().split()))
s = input()
ans = 0
l, r = 0, 1
while r < n:
    if s[l] == s[r]:
        r += 1
    else:
        if r - l <= k:
            ans += sum(arr[l:r])
        else:
            heap = []
            for i in range(l, r):
                heapq.heappush(heap, arr[i])
                while len(heap) > k:
                    heapq.heappop(heap)
            ans += sum(heap)
        l = r
        r = l + 1
if r - l <= k:
    ans += sum(arr[l:r])
else:
    heap = []
    for i in range(l, r):
        heapq.heappush(heap, arr[i])
        while len(heap) > k:
            heapq.heappop(heap)
    ans += sum(heap)
print(ans)
