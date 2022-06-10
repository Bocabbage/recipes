#include "EchoServer.hpp"
#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include <sys/epoll.h>
#include <unistd.h>      // close
#include <strings.h>    // bzero
#include <iostream>
#include <cstring>      // strerror


EchoServer::EchoServer(
    std::string ip,
    short port
):
ip_(ip),
port_(port),
epollFd_(-1),
listenFd_(-1),
threadNums_(4)  // hard-code to 4 threads
{
    // Initialize the threadPool
    for(int i = 0; i < threadNums_; ++i)
        threadPool_.push_back(std::thread(&message_handler, this));
}

EchoServer::~EchoServer()
{

}

/*
    Function start:
    1. Create listenfd(non-block) and bind-listen
    2. Create epollfd
    3. Add listenfd to epoll-wait-sys
*/
bool EchoServer::start()
{
    /* create listen-fd */
    listenFd_ = ::socket(
        AF_INET, 
        SOCK_STREAM | SOCK_NONBLOCK | SOCK_CLOEXEC,
        0
    );
    if(listenFd_ < 1)
    {
        std::cout << "listenFd_ creation failed." << std::endl;
        return false;
        // exit(-1);
    }

    struct sockaddr_in servAddr;
    bzero((void *)&servAddr, sizeof(servAddr));
    servAddr.sin_family = AF_INET;
    servAddr.sin_port = htons(port_);
    inet_pton(AF_INET, ip_.c_str(), &(servAddr.sin_addr));

    if(::bind(listenFd_, (struct sockaddr *)&servAddr, sizeof(servAddr)) < 0)
    {
        std::cout << "listenFd_ bind failed." << std::endl;           
        return false;
        // exit(-1);
    }

    if(::listen(listenFd_, 50) < -1)
    {
        std::cout << "listenFd_ listen failed." << std::endl;
        return false;
        // exit(-1);
    }

    /* create epollFd_ */
    epollFd_ = ::epoll_create1(EPOLL_CLOEXEC);
    
    // Add listenFd_ 
    ::epoll_event listenEvent;
    listenEvent.events = EPOLLIN;
    listenEvent.data.fd = listenFd_;
    if(::epoll_ctl(epollFd_, EPOLL_CTL_ADD, listenFd_, &listenEvent) < 0)
    {
        std::cout << "Add listenFd_ to EPOLL failed." << std::endl;           
        return false;
        // exit(-1);
    }

    std::cout << "server start" << std::endl;
    return main_loop();

}

bool EchoServer::main_loop()
{
    std::vector<epoll_event> epollEvents(50);
    while(true)
    {
        int n = ::epoll_wait(epollFd_, &(*epollEvents.begin()), 51, 10);
        if(n < 0)
        {
            // std::cout << "epoll_wait: " << n << " failed." << std::endl;
            // //return false;
            // continue;

            int error = 0;
            socklen_t errlen = sizeof(error);
            if (getsockopt(epollFd_, SOL_SOCKET, SO_ERROR, (void *)&error, &errlen) == 0)
            {
                std::cout << "error = " << std::strerror(error) << " %s\n";
            }
            continue;

        }
        for(int i = 0; i < n; ++i)
        {
            if(epollEvents[i].data.fd == listenFd_)
                accept_handler();
            else
            {
                {
                    std::lock_guard<std::mutex> guard(taskQueueLk_);
                    taskQueue_.push(epollEvents[i].data.fd);
                }
                taskCond_.notify_one();
            }
        }
    }

    return true;
}

void EchoServer::accept_handler()
{
    struct sockaddr_in newAddr;
    socklen_t addrlen;
    int newFd = ::accept4(
        listenFd_,
        (struct sockaddr *)&newAddr,
        &addrlen,
        SOCK_NONBLOCK
    );

    ::epoll_event clientEvent;
    clientEvent.events = EPOLLIN | EPOLLRDHUP | EPOLLET; // ET-mode
    clientEvent.data.fd = newFd;
    if(::epoll_ctl(epollFd_, EPOLL_CTL_ADD, newFd, &clientEvent) < 0)
    {
        std::cout << "Add clientFd_ to EPOLL failed." << std::endl;           
        // exit(-1);
    }

}

void EchoServer::message_handler(EchoServer *pEchoServer)
{
    while(true)
    {
        // Take task from the taskQueue to process
        // If taskQueue empty, wait on the taskCond_
        std::unique_lock<std::mutex> lk(pEchoServer->taskQueueLk_);
        pEchoServer->taskCond_.wait(
            lk,
            [echo = pEchoServer] { return !(echo->taskQueue_.empty()); }
            // Buggy: [tkq = pEchoServer->taskQueue_] { return !(tkq.empty()); }
        );

        int taskFd = pEchoServer->taskQueue_.front();
        pEchoServer->taskQueue_.pop();
        lk.unlock();

        /* the ECHO process */
        std::string recvMsg;
        char buff[256];

        // receive message
        bool errorOccur = false;
        while(true)
        {
            bzero(&buff, 256);
            int recvN = ::recv(taskFd, &buff, 256, 0);
            if(recvN < 0)
            {
                if(errno == EWOULDBLOCK)
                    break;  // have received all the message
                else
                {
                    std::cout << "recv-function error.\n" << std::endl;
                    errorOccur = true;
                    break;
                }
            }
            else if(recvN == 0)
            {
                // the connection would be shut-down
                pEchoServer->shutdown_handler(taskFd);
                errorOccur = true;
                break;
            }

            recvMsg += buff;
        }

        // send message
        if(errorOccur)
            continue;
        
        // Add server-tag
        recvMsg = "[Echo]" + recvMsg;

        while(true)
        {
            int sendN = ::send(taskFd, recvMsg.c_str(), recvMsg.length(), 0);
            if(sendN < 0)
            {
                if(errno == EWOULDBLOCK)
                {
                    std::this_thread::sleep_for(std::chrono::milliseconds(10));  
                    continue;
                }
                else
                {
                    std::cout << "send-function error.\n" << std::endl;
                    pEchoServer->shutdown_handler(taskFd);
                    break;
                }
            }

            recvMsg.erase(0, sendN);
            if(recvMsg.empty())
                break;
        }

    }
}

void EchoServer::shutdown_handler(int fd)
{
    if(::epoll_ctl(epollFd_, EPOLL_CTL_DEL, fd, NULL) < 0)
    {
        std::cout << "close fd: " << fd << " error." << std::endl;
        return;
    }

    std::cout << "close fd: " << fd << std::endl;
    ::close(fd);
}