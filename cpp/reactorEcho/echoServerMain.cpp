#include "EchoServer.hpp"

int main()
{
    EchoServer server("0.0.0.0", 1234);
    server.start();
}