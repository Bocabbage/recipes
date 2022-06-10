// Update date: 2022/02/16
// Author: Zhuofan Zhang
#ifndef ECHOSERVER_HPP
#define ECHOSERVER_HPP
#include <string>
#include <thread>
#include <vector>
#include <queue>
#include <mutex>
#include <condition_variable>

class EchoServer
{
public:
    EchoServer(std::string ip, short port);
    ~EchoServer();
    // Non-copyable
    EchoServer(EchoServer &) = delete;
    EchoServer& operator=(const EchoServer &) = delete;

    bool start();

private:
    int epollFd_;
    int listenFd_;
    std::string ip_;    // local ip-address
    short port_;          // local port

    // threadpool-realted
    const int threadNums_;
    std::vector<std::thread> threadPool_;

    // task-queue
    std::queue<int> taskQueue_;

    // mutex-lock and cond
    mutable std::mutex taskQueueLk_;
    mutable std::mutex threadPoolLk_;
    std::condition_variable taskCond_;

    bool main_loop();
    void accept_handler();
    static void message_handler(EchoServer *);
    void shutdown_handler(int);

};

#endif