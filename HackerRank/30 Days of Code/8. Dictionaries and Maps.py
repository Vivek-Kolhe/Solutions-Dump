import sys

n = int(input())
phonebook = {}
for _ in range(n):
    contact = input().split()
    phonebook[contact[0]] = contact[1]

line = sys.stdin.readlines()
for i in line:
    name = i.strip()
    if name in phonebook:
        print(name + "=" + str(phonebook[name]))
    else:
        print("Not found")
