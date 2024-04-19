#include <future>
#include <chrono>
#include <iostream>
using namespace std;

int main() {
    // the task is executed on a different thread, potentially by creating and launching it first
    auto fu = std::async(std::launch::async, [](){ std::cout << "Hello from async-baby!" << std::endl; });
    // the task is executed on the calling thread the first time its result is requested (lazy evaluation)
    auto fu2 = std::async(std::launch::deferred, [](){ std::cout << "Hello from deferred-baby!" << std::endl; });
    
    // Main thread sleep
    std::this_thread::sleep_for(std::chrono::seconds(5));
    std::cout << "Hello main!~" << std::endl;

    fu.wait();
    fu2.wait(); // deferred-baby will print here
}