# cook your dish here
r, o, c = map(int, input().split())
can_make = (20 - o) * 6 * 6
if r - c < can_make:
    print("yes")
else:
    print("no")
