#include <iostream>
using namespace std;

template<typename T>
auto foo(T y)
{
    return y + 2;
}


int main() {
    char z = 'x';
    auto myLambda = [y = z + 2](int x) {
        std::cout << x << '-' << y << std::endl;
    };

    myLambda(12);
    foo(z+2);
}