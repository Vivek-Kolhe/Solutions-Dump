#include <iostream>
#include <cstdio>
using namespace std;

int max_of_four(int a, int b, int c, int d)
{
    int arr[4] = {a, b, c, d};
    int max_num = arr[0];
    for(int i = 0; i < 4; i++)
    {
        if(arr[i] > max_num)
        {
            max_num = arr[i];
        }
    }
    return max_num;
}

int main() {
    int a, b, c, d;
    scanf("%d %d %d %d", &a, &b, &c, &d);
    int ans = max_of_four(a, b, c, d);
    printf("%d", ans);
    
    return 0;
}
