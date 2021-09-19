#include <stdio.h>
#include <iostream>
#include <algorithm>
#include <string.h>
using namespace std;
const int maxn = 100005;
bool num[maxn];
int main()
{
    int N;
    scanf("%d", &N);
    memset(num, false, sizeof(num));
    int max_num = 0;
    int flag_num = 0;
    for(int i = 0; i < N * 2; i++)
    {
        int flag;
        scanf("%d", &flag);
        if(num[flag] == false)
        {
            num[flag] = true;
            flag_num++;
        }
        else
        {
            num[flag] = false;
            flag_num--;
        }
        max_num=max(max_num, flag_num);
    }
    printf("%d\n", max_num);
    return 0;
}
