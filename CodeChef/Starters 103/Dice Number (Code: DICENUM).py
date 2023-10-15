# cook your dish here
for _ in range(int(input())):
    arr = list(map(int, input().split()))
    alice = sorted(arr[:3])[::-1]
    bob = sorted(arr[3:])[::-1]
    alice, bob = list(map(str, alice)), list(map(str, bob))
    alice, bob = int("".join(alice)), int("".join(bob))
    if alice > bob: print("Alice")
    elif alice < bob: print("Bob")
    else: print("Tie")
