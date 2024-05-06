cmake_minimum_required(VERSION 3.5.0)
project(cmake-layout-demo VERSION 0.1.0 LANGUAGES C CXX)

# include(CTest)
# enable_testing()
add_subdirectory(src/)
include_directories(include/)

add_executable(cmake-layout-demo main.cpp)
target_link_libraries(cmake-layout-demo demolib)

set(CPACK_PROJECT_NAME ${PROJECT_NAME})
set(CPACK_PROJECT_VERSION ${PROJECT_VERSION})
include(CPack)
