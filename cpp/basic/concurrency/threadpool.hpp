#include <iostream>
#include <thread>
#include <mutex>
#include <chrono>
#include <time.h>
#include <vector>
#include <queue>
#include <future>
#include <mutex>
#include <queue>
#include <functional>
#include <thread>
#include <utility>
#include <vector>
#include <condition_variable>
#include <string>
#include <shared_mutex>
using namespace std;

// 涉及内容
// 1. 类型推导（auto/decltype/lambda）
// 2. concurrency 特性（std::future, std::packaged_task）
// 3. 调用对象封装
// 4. 移动语义 + 完美转发


// 线程安全队列
template<typename T>
struct safe_queue {
    queue<T>que;
    shared_mutex _m;

    bool empty() {
        shared_lock<shared_mutex>lc(_m); // 读锁
        return que.empty();
        // 离开作用域，lc 析构释放_m
    }
    auto size() {
        shared_lock<shared_mutex>lc(_m); // 读锁
        return que.size();
    }
    void push(T& t) {
        unique_lock<shared_mutex> lc(_m); // 写锁
        que.push(t);
    }
    bool pop(T& t) {
        unique_lock<shared_mutex> lc(_m); // 写锁
        if (que.empty())return false;
        t = move(que.front());
        que.pop();
        return true;
    }
};


class ThreadPool {
private:
    // 内嵌 class 定义
    class worker {
    public:
        ThreadPool* pool;
        worker(ThreadPool* _pool) : pool{ _pool } {}
        void operator()() {
            while (!pool->is_shut_down) {
                {
                    unique_lock<mutex> lock(pool->_m);
                    pool->cv.wait(lock, [this]() {
                        // 在线程池关闭或有新的活儿的时候唤醒 worker
                        return this->pool->is_shut_down ||
                            !this->pool->que.empty();
                        });
                }
                function<void()>func;
                bool flag = pool->que.pop(func);
                if (flag) {
                    func();
                }
            }
        }
    };
public:
    bool is_shut_down;
    safe_queue<std::function<void()>> que;
    vector<std::thread>threads;
    mutex _m;
    condition_variable cv;
    ThreadPool(int n) : threads(n), is_shut_down{ false } {
        for (auto& t : threads)t = thread{ worker(this) };
    }
    ThreadPool(const ThreadPool&) = delete;
    ThreadPool(ThreadPool&&) = delete;
    ThreadPool& operator=(const ThreadPool&) = delete;
    ThreadPool& operator=(ThreadPool&&) = delete;

    // 向线程池提交任务
    template <typename F, typename... Args>
    auto submit(F&& f, Args &&...args) -> std::future<decltype(f(args...))> {
        // 万能引用
        
        // 尾置返回类型：auto funcname(args) -> return_type
        // 1. 支持返回类型依赖于参数类型的情况
        // 2. 支持 lambda 表达式和 decltype

        // lambda 表达式对象：捕获&f与args转化为无参函数调用
        function<decltype(f(args...))()> func = [&f, args...]() {return f(args...); };
        // std::packaged_task 是 C++11 中引入的一个模板类，它将一个可调用对象包装起来，使其可以异步执行。
        // 它的模板参数是一个函数签名，这意味着你可以将任何可以调用的对象（比如函数指针、lambda 表达式、bind 表达式或其他函数对象）包装到 std::packaged_task 中，只要这些对象的调用签名与模板参数匹配。
        // 当 std::packaged_task 对象的 operator() 被调用时，它会调用相关联的可调用对象，并将返回值（如果有的话）存储在一个共享状态中。
        // 这个共享状态可以通过一个与 std::packaged_task 相关联的 std::future 对象来访问。
        // 这使得 std::packaged_task 成为创建异步代码的一种方式，因为你可以在一个线程中执行 std::packaged_task，并在另一个线程中等待结果。
        auto task_ptr = std::make_shared<std::packaged_task<decltype(f(args...))()>>(func);
        std::function<void()> warpper_func = [task_ptr]() {
            (*task_ptr)(); // 把实际可调用对象塞进queue里，调用结果可获取时，对应的future那就可以get到
        };
        que.push(warpper_func);
        cv.notify_one(); // 通知消费者
        return task_ptr->get_future();
        // std::future 对象用于提供访问异步操作结果的机制
    }

    ~ThreadPool() {
        auto f = submit([]() {});
        f.get();
        is_shut_down = true;
        cv.notify_all(); // 通知，唤醒所有工作线程
        for (auto& t : threads) {
            if (t.joinable()) t.join();
        }
    }
};
