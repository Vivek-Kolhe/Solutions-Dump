from itertools import permutations
permutations_ = list(permutations(["A", "B", "C"]))
permutations_ = ["".join(item) for item in permutations_]
s = input().upper()
flag = False
for i in range(len(permutations_)):
    if permutations_[i] in s:
        flag = True
        break
if flag:
    print("Yes")
else:
    print("No")
