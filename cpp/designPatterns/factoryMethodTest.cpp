#include "factoryMethod.hpp"
#include <iostream>
#include <memory>

using std::cout;
using std::endl;
using std::unique_ptr;

int main()
{
    // Use base-product-class pointer:
    // decoupling from concrete-product-class obj
    unique_ptr<Product> ptrA = AFactory().createMethod();
    unique_ptr<Product> ptrB = BFactory().createMethod();

    cout << "Product A type: " << ptrA->type() << endl;
    cout << "Product B type: " << ptrB->type() << endl;
}