#include "test_lib.h"
#include <string.h>
#include <errno.h>

void
hello(const char* msg)
{
    if(strlen(msg) == 0)
    {
        errno = EINVAL;
        return;
    }

    printf("Hello, my friend %s", msg);
}