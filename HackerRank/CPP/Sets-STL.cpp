#include <cmath>
#include <cstdio>
#include <vector>
#include <iostream>
#include <set>
#include <algorithm>
using namespace std;

int main()
{
    int n;
    cin >> n;
    set<int> s;
    for(int i = 0; i < n; i++)
    {
        int type, num;
        cin >> type >> num;
        if(type == 1)
        {
            s.insert(num);
        }
        else if(type == 2)
        {
            s.erase(num);
        }
        else if(type == 3)
        {
            set<int>::iterator it = s.find(num);
            if(it != s.end())
            {
                cout << "Yes" << endl;
            }
            else
            {
                cout << "No" << endl;
            }
        }
    }
    return 0;
}
