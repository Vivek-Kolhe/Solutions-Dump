#include <iostream>
#include <sstream>
using namespace std;

class Student
{
    int s_age, s_standard;
    string f_name, l_name;
    public:
        void set_age(int age)
        {
            s_age = age;
        }
        void set_standard(int standard)
        {
            s_standard = standard;
        }
        void set_first_name(string first_name)
        {
            f_name = first_name;
        }
        void set_last_name(string last_name)
        {
            l_name = last_name;
        }
        
        int get_age()
        {
            return s_age;
        }
        int get_standard()
        {
            return s_standard;
        }
        string get_first_name()
        {
            return f_name;
        }
        string get_last_name()
        {
            return l_name;
        }
        
        string to_string()
        {
            string s;
            stringstream ss;
            ss << s_age << "," << f_name << "," << l_name << "," << s_standard;
            s = ss.str();
            return s;
        }
};

int main() {
    int age, standard;
    string first_name, last_name;
    
    cin >> age >> first_name >> last_name >> standard;
    
    Student st;
    st.set_age(age);
    st.set_standard(standard);
    st.set_first_name(first_name);
    st.set_last_name(last_name);
    
    cout << st.get_age() << "\n";
    cout << st.get_last_name() << ", " << st.get_first_name() << "\n";
    cout << st.get_standard() << "\n";
    cout << "\n";
    cout << st.to_string();
    
    return 0;
}
