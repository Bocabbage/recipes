cmake_minimum_required(VERSION 3.16)
project(reactorEcho)

set(SRC_LIST EchoServer.cpp)
add_executable(reactorEchoServer echoServerMain.cpp ${SRC_LIST})

set(THREADS_PREFER_PTHREAD_FLAG ON)
find_package(Threads REQUIRED)
target_link_libraries(reactorEchoServer PRIVATE Threads::Threads)

