#include <dlfcn.h>
#include <iostream>
// #include "mocklib_cmode.hpp"

class DynamicLibHelper
{
    void *_handler;
public:
    DynamicLibHelper(const std::string& path) 
    {
        _handler = dlopen(path.c_str(), RTLD_LAZY);
        if(!_handler)
        {
            std::cerr << "Can't open library: " << dlerror() << std::endl;
            exit(-1);
        }
        dlerror();
        std::cout << "Init dl-handle." << std::endl;
    }

    ~DynamicLibHelper()
    {
        std::cout << "Close dl-handle." << std::endl;
        if(_handler)
        {
            dlclose(_handler);
        }
    }

    template<typename FP>
    auto load_func(const std::string& funcname) -> FP
    {
        auto func = (FP) dlsym(_handler, funcname.c_str());
        auto error = dlerror();
        if(error)
        {
            std::cerr << "Can't load symbol hello: " << error << std::endl;
            return nullptr;
        }
        return func;
    }

    DynamicLibHelper(DynamicLibHelper&)=delete;
    DynamicLibHelper(DynamicLibHelper&&)=delete;
    DynamicLibHelper& operator=(const DynamicLibHelper&)=delete;
    DynamicLibHelper& operator=(DynamicLibHelper&&)=delete;
};

int main() 
{
    using HelloFunc = void (*)(const char*); // define function type

    // Lazy load .so file
    DynamicLibHelper dh("./mocklib_c.so");

    // Load symbol
    auto hello = dh.load_func<HelloFunc>("hello");
    const char *error = dlerror();
    if(error)
    {
        std::cerr << "Can't load symbol hello: " << error << std::endl;
        return -1;
    }

    // Call function
    hello("Bocabbage");

    return 0;
}