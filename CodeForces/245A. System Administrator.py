n = int(input())
recieve_a, lost_a = 0, 0
recieve_b, lost_b = 0, 0
for _ in range(n):
    server, recieved, lost = map(int, input().split())
    if server == 1:
        recieve_a += recieved
        lost_a += lost
    else:
        recieve_b += recieved
        lost_b += lost
if recieve_a >= (recieve_a + lost_a) / 2:
    print("LIVE")
else:
    print("DEAD")
if recieve_b >= (recieve_b + lost_b) / 2:
    print("LIVE")
else:
    print("DEAD")
