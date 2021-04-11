#include <cmath>
#include <cstdio>
#include <vector>
#include <iostream>
#include <algorithm>
#include <cassert>
using namespace std;

/*
without start_up() code will run for only few test cases
function just only fasts up the cout
and #define changes endl to \n which prevents flushing buffer.
(found this in the problem comments while i was looking for what did i do wrong.)

int start_up() {
    ios_base::sync_with_stdio(false);
    cin.tie(NULL);
    return 0;
}

int static r = start_up();

#define endl '\n'
*/

template <class T>
class AddElements
{
    T element;
    public:
        AddElements(T arg_1)
        {
            element = arg_1;
        }
        T add(T arg_2)
        {
            return(element + arg_2);
        }
};
template <>
class AddElements <string>
{
    string element;
    public:
        AddElements(string arg_1)
        {
            element = arg_1;
        }
        string concatenate(string arg_2)
        {
            return(element + arg_2);
        }
};

int main () {
  int n,i;
  cin >> n;
  for(i=0;i<n;i++) {
    string type;
    cin >> type;
    if(type=="float") {
        double element1,element2;
        cin >> element1 >> element2;
        AddElements<double> myfloat (element1);
        cout << myfloat.add(element2) << endl;
    }
    else if(type == "int") {
        int element1, element2;
        cin >> element1 >> element2;
        AddElements<int> myint (element1);
        cout << myint.add(element2) << endl;
    }
    else if(type == "string") {
        string element1, element2;
        cin >> element1 >> element2;
        AddElements<string> mystring (element1);
        cout << mystring.concatenate(element2) << endl;
    }
  }
  return 0;
}
