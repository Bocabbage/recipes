#include <iostream>
#include <future>
#include <thread>

int main() {
    // 创建一个std::packaged_task，包装一个lambda表达式
    std::packaged_task<int()> task([](){
        return 7; // 假设这是一个复杂的计算
    });

    // 获取与std::packaged_task相关联的std::future对象
    std::future<int> result = task.get_future();

    // 在一个新线程上执行task
    std::thread t(std::move(task));

    // 在主线程上等待结果
    std::cout << "Waiting for result...\n";
    std::cout << "Result is: " << result.get() << std::endl;

    // 等待任务线程结束
    t.join();

    return 0;
}
