#include "mocklib_cmode.hpp"

extern "C" {
    void 
    hello(const char *msg) {
        std::cout << "Hello, my friend " << msg << std::endl;
    }
}