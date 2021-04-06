#include <iostream>
#include <string>
using namespace std;

int main()
{
    string a, b;
    string temp;
    cin >> a;
    cin >> b;
    
    cout << a.size() << " " << b.size() << endl;
    cout << a + b << endl;
    
    temp = a;
    a[0] = b[0];
    b[0] = temp[0];
    cout << a << " " << b << endl;
    return 0;
}
