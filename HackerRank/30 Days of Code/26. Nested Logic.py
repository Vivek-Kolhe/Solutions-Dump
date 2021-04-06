r = list(map(int, input().split()))
e = list(map(int, input().split()))

if r[2] < e[2]:
    print(0)
elif r[0] <= e[0] and r[1] <= e[1] and r[2] <= e[2]:
    print(0)
elif r[0] > e[0] and r[1] == e[1] and r[2] == e[2]:
    print(15 * (r[0] - e[0]))
elif r[1] > e[1] and r[2] == e[2]:
    print(500 * (r[1] - e[1]))
elif r[2] > e[2]:
    print(10000)
