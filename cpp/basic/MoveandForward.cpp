#include <vector>
#include <iostream>

using namespace std;

class TestClass
{
public:
    TestClass()
    { cout << "TestClass-defaultConstructor called." << endl; }

    TestClass(const TestClass& tc)
    { cout << "TestClass-copyConstructor called." << endl; }

    TestClass(TestClass&& tc)
    { cout << "TestClass-moveConstructor called." << endl; }

    TestClass& operator=(const TestClass& tc)
    { cout << "TestClass-copyAssign called." << endl; return *this; }

    TestClass& operator=(TestClass&& tc)
    { cout << "TestClass-moveAssign called." << endl; return *this; }

};

int main()
{
    vector<TestClass> v;
    //!
    v.reserve(10);
    
    cout << "v.emplace_back(): " << endl;
    v.emplace_back();   // call default-constructor

    cout << "v.emplace_back(TestClass()): " << endl;
    v.emplace_back(TestClass());

    // std::move() here makes no sense: TestClass() is rvalue
    cout << "v.emplace_back(std::move(TestClass())): " << endl;
    v.emplace_back(std::move(TestClass()));

    // push_back(T&&) called emplace_back(T&&) used std::move()
    cout << "v.push_back(TestClass()): " << endl;
    v.push_back(TestClass());

    // std::move() here makes no sense: TestClass() is rvalue
    cout << "v.push_back(std::move(TestClass())): " << endl;
    v.push_back(std::move(TestClass()));

    cout << "v.back() = TestClass(): " << endl;
    v.back() = TestClass();

    cout << "TestClass nobj = TestClass(): " << endl;
    TestClass nobj = TestClass();

}