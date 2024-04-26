# Compile mocklib_cmode
g++ -shared -fPIC -o mocklib_c.so mocklib_cmode.cpp

# Compile mocklib_cmode user code
g++ -std=c++14 -o cmode_main.out cmode_main.cpp -ldl
# Check file type
file cmode_main.out