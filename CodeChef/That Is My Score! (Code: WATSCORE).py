for _ in range(int(input())):
    max_ = {
        1 : 0,
        2 : 0,
        3 : 0,
        4 : 0,
        5 : 0,
        6 : 0,
        7 : 0,
        8 : 0
    }
    for i in range(int(input())):
        p, marks = map(int, input().split())
        if p < 9:
            if marks > max_[p]:
                max_[p] = marks
    print(sum(list(max_.values())))
