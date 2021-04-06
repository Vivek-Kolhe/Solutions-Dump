#include <iostream>
using namespace std;

class Rectangle
{
    public:
        int breadth, length;
        void read_input()
        {
            cin >> breadth >> length;
        }
        void display()
        {
            cout << breadth << " " << length << endl;
        }
};

class RectangleArea: public Rectangle
{
    public:
        void display()
        {
            cout << Rectangle::length * Rectangle::breadth << endl;
        }
};

int main()
{
    RectangleArea r_area;
    r_area.read_input();
    r_area.Rectangle::display();
    r_area.display();
    return 0;
}
