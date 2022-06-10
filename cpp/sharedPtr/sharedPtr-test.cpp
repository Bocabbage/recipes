#include "sharedPtr.hpp"
#include <iostream>

using std::cout;
using std::endl;

int main()
{
    sharedPtr<int> sp(new int(10));
    cout << "sp value: " << *sp << endl;
    cout << "user_count: " << sp.user_count() << endl;

    auto sp2(sp);
    *sp2 = 15;
    cout << "sp value: " << *sp << endl;
    cout << "sp2 value: " << *sp2 << endl;
    cout << "sp user_count: " << sp.user_count() << endl;
    cout << "sp2 user_count: " << sp2.user_count() << endl;

    auto sp3 = sp;
    *sp3 = 20;
    cout << "sp value: " << *sp << endl;
    cout << "sp2 value: " << *sp2 << endl;
    cout << "sp user_count: " << sp.user_count() << endl;
    cout << "sp2 user_count: " << sp2.user_count() << endl;
}