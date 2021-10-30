n = int(input())
string = input()
if n % 4 != 0:
  print("===")
else:
  per = n // 4
  count_a, count_c, count_g, count_t = string.count("A"), string.count("C"), string.count("G"), string.count("T")
  if count_a > per or count_c > per or count_g > per or count_t > per:
    print("===")
  else:
    remain_a, remain_c, remain_g, remain_t = per - count_a, per - count_c, per - count_g, per - count_t
    string = string.replace("?", "A", remain_a)
    string = string.replace("?", "C", remain_c)
    string = string.replace("?", "G", remain_g)
    string = string.replace("?", "T", remain_t)
    print(string)
