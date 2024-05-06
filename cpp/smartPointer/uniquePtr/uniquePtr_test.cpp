#include <iostream>
#include "uniquePtr.hpp"

using std::cout;
using std::endl;
using namespace demosp;

struct TestStruct {
    int a;
    double b;
    char c;

    TestStruct(int _a, double _b, char _c):
    a(_a), b(_b), c(_c)
    {  }
};

int main() {
    auto upt = make_unique<TestStruct>(3, 4.0, 'c');
    // test: operator->
    cout << upt->a << endl;
    // test: operator*
    cout << (*upt).b << endl;
    // test operator.get
    cout << upt.get()->c << endl;

    // test: reset
    upt.reset(new TestStruct(4, 5.0, 'd'));
    cout << upt->a << endl;

    // test: swap
    // (1) self-swap
    upt.swap(upt);
    cout << upt->a << endl;
    // (2) normal-swap
    auto upt2 = make_unique<TestStruct>(5, 6.0, 'e');
    upt.swap(upt2);
    cout << upt->a << endl;
    cout << upt2->a << endl;

    // test: release
    upt2.release();
    if(!upt2.get())
        cout << "upt2 cleaned" << endl;
}