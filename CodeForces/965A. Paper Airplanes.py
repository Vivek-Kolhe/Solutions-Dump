k, n, s, p = map(int, input().split())
per_person = n // s if n % s == 0 else (n // s) + 1
tot_sheets = k * per_person
print((tot_sheets // p) + 1 if tot_sheets % p != 0 else tot_sheets // p)
