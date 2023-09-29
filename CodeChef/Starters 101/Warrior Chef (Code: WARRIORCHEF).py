# cook your dish here
def check_hp(arr, hp, resis):
    for i in range(len(arr)):
        if arr[i] > resis:
            hp -= arr[i]
    if hp > 0:
        return True
    return False

for _ in range(int(input())):
    n, h = map(int, input().split())
    arr = list(map(int, input().split()))
    if h > sum(arr):
        print(0)
    else:
        ans = 0
        low, high = 0, max(arr) + 1
        while low <= high:
            mid = (low + high) // 2
            if check_hp(arr, h, mid):
                ans = mid
                high = mid - 1
            else:
                low = mid + 1
        print(ans)
