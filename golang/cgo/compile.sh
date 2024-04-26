# Compile mocklib_cmode
gcc -shared -fPIC -o test_lib.so test_lib.c

# Compile golang program
go build -o main.out # Can't use: go build main.go
# OR: 'go build .' to use default output name