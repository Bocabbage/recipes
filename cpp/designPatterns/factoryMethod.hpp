#pragma once
#include <string>
#include <memory>
using std::string;
using std::unique_ptr;

// product base-class
class Product
{
public:
    Product(const string& type):
        type_(type)
    {  }

    Product(string&& type):
        type_(std::move(type))
    {  }

    string type() const { return type_; }

private:
    string type_;
};

// Abstract creator base-class
class Creator
{
public:
    virtual unique_ptr<Product> createMethod() = 0;
};

// -----------------------------------------------

// Concrete Product A
class ProductA: public Product
{
public:
    ProductA():
        Product("typeA")
    {  }

};

// Concrete Product B
class ProductB: public Product
{
public:
    ProductB():
        Product("typeB")
    {  }
};

// Concrete CreatorA
class AFactory: public Creator
{
public:
    AFactory() = default;
    ~AFactory() = default;

    unique_ptr<Product> createMethod()
    {
        unique_ptr<Product> productPtr(new ProductA);
        return productPtr;
    }

};

// Concrete CreatorB
class BFactory: public Creator
{
public:
    BFactory() = default;
    ~BFactory() = default;

    unique_ptr<Product> createMethod()
    {
        unique_ptr<Product> productPtr(new ProductB);
        return productPtr;
    }
 
};
