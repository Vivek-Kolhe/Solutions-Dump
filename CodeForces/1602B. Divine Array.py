for _ in range(int(input())):
    n = int(input())
    arr = list(map(int, input().split()))
    queries = []
    max_ = 0
    for _ in range(int(input())):
        x, k = map(int, input().split())
        if max_ < k:
            max_ = k
        queries.append((x, k))
    new_arr = []
    new_arr.append(arr)
    for i in range(max_):
        temp = [0] * n
        for j in range(len(arr)):
            temp[j] = new_arr[-1].count(new_arr[-1][j])
        if temp == new_arr[-1]:
            break
        new_arr.append(temp)
    for _ in range(len(queries)):
        try:
            print(new_arr[queries[_][-1]][queries[_][0]-1])
        except:
            print(new_arr[-1][queries[_][0]-1])
